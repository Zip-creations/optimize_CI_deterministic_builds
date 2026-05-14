package main

import "fmt"
import "os"
import "os/exec"
import "encoding/xml"
import "path/filepath"
import "bytes"
import dt "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/pipeline_observer/datatypes"

func main() {
	// Read all existing tests from the user-configured script
	output, err := RunTestDiscoveryScript("examples/sample_find.sh")  // TODO: Read from config.json
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output, "\n")  // Debug

	// Read all tests in the JUnit XML output of the last run (if existing)
	allSuites, err := ReadTestSuites("./examples/jUnit_XML")  // TODO: Read from config.json
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(allSuites, "\n")  // Debug

	report := CreateReport("Test Report", allSuites)
	WriteXMLToFile(report, "./out/report.xml")
	fmt.Println("Successfully created report: \n", report)  // Debug
}

func RunTestDiscoveryScript(path string) (dt.Testsuite, error) {
	cmd := exec.Command("bash", path)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return dt.Testsuite{}, fmt.Errorf("Error executing test discovery script: %w\n%s", err, stderr.String())
	}
	return XMLtoTestSuite([]byte(out.String()), &dt.Testsuite{})
}

func ReadTestSuites(path string) (dt.Testsuites, error) {
	var allSuites dt.Testsuites
	content, err := os.ReadDir(path)
	if err != nil {
		return allSuites, fmt.Errorf("Error reading directory:\n %s\n %w", path, err)
	}
	content = filterForXML(content)
	for _, entry := range content {
		filePath := filepath.Join(path, entry.Name())
		testSuit, err := ReadTestSuite(filePath)
		if err != nil {
			fmt.Println(err)  // TODO: log error somehow
			continue  // If one file is broken: skip and continue with the others
		}
		allSuites.Testsuites = append(allSuites.Testsuites, testSuit)
	}
	return allSuites, nil
}

func ReadTestSuite(filePath string) (dt.Testsuite, error) {
	var testsuite dt.Testsuite
	data, err := os.ReadFile(filePath)
	if err != nil {
		return testsuite, fmt.Errorf("Error while reading file:\n %s\n %w", filePath, err)
	}
	return XMLtoTestSuite(data, &testsuite)
}

func XMLtoTestSuite(data []byte, suite *dt.Testsuite) (dt.Testsuite, error) {
	err := xml.Unmarshal(data, suite)
	if err != nil {
		return *suite, fmt.Errorf("Error while unmarshalling TestSuite XML:\n %w", err)
	}
	return *suite, err
}

func filterForXML(files []os.DirEntry) []os.DirEntry {
	var xmlFiles []os.DirEntry
	for _, file := range files {
		if file.Name()[len(file.Name())-4:] == ".xml" {
			xmlFiles = append(xmlFiles, file)
		}
	}
	return xmlFiles
}

func CreateReport(name string, testSuites dt.Testsuites) Report {
	var totalRun, totalFailed, totalSkipped int = 0, 0, 0
	allSuites := []TestSuiteReport{}
	for _, testsuite := range testSuites.Testsuites {
		allSuites = append(allSuites, CreateTestSuiteReport(testsuite, &totalRun, &totalFailed, &totalSkipped))
	}
	return Report{
		Name: name,
		TestsTotal: totalRun + totalSkipped,
		TestsRun: totalRun,
		// TestsFailed: totalFailed,
		TestsSkipped: totalSkipped,
		TestSuites: allSuites,
	}
}

func CreateTestSuiteReport(testsuite dt.Testsuite, totalRun *int, totalFailed *int, totalSkipped *int) TestSuiteReport {
	var testCases []TestCaseReport
	var totalRunSuite, totalFailedSuite, totalSkippedSuite int = 0, 0, 0
	for _, testcase := range testsuite.Testcases {
		var result TestStatus
		if testcase.IsSkipped() {
			*totalSkipped++
			totalSkippedSuite++
			result = StatusSkipped
		} else if testcase.HasFailed() {
			*totalFailed++
			totalFailedSuite++
			result = StatusFailed
		} else {
			result = StatusPassed
		}
		*totalRun++
		totalRunSuite++
		testCases = append(testCases, TestCaseReport{
			Name:   testcase.Name,
			Result: result,
		})
	}
	return TestSuiteReport{
		Name:          testsuite.Name,
		Timestamp: testsuite.Timestamp,
		TestsTotal: totalRunSuite + totalSkippedSuite,
		TestsRun: totalRunSuite,
		// TestsFailedSuite: totalFailedSuite,
		TestsSkipped: totalSkippedSuite,
		TestCases:     testCases,
	}
}

func WriteXMLToFile(report Report, filePath string) error {
	data, err := xml.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("Error while marshalling Report XML:\n %w", err)
	}
	return os.WriteFile(filePath, data, 0644)
}
