package main

import "fmt"
import "os"
import "encoding/xml"
import "path/filepath"

func main() {
	path := "./examples"  // TODO: ask path on first start of tool. NiceToHave: Make it configurable
	content, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading files: ", err)
		return
	}
	content = filterForXML(content)
	var allSuites []TestSuiteReport
	var totalFailed, totalSkipped int = 0, 0
	for _, entry := range content {
		filePath := filepath.Join(path, entry.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error unmarshaling XML: ", err)
			continue
		}

		var testsuite Testsuite
		err = xml.Unmarshal(data, &testsuite)
		if err != nil {
			fmt.Println("Error unmarshaling XML: ", err)
			continue
		}
		allSuites = append(allSuites, CreateTestSuiteReport(testsuite, &totalFailed, &totalSkipped))
		for _, testcase := range testsuite.Testcases {
			fmt.Println("Found:", testcase.IsSkipped())  // Debug
		}
	}
	report := CreateReport("Test Report", totalFailed, totalSkipped, allSuites)
	WriteXMLToFile(report, "./out/report.xml")
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

func CreateReport(name string, totalFailed, totalSkipped int, testSuites []TestSuiteReport) Report {
	return Report{
		Name: name,
		TotalFailed: totalFailed,
		TotalSkipped: totalSkipped,
		TestSuites: testSuites,
	}
}

func CreateTestSuiteReport(testsuite Testsuite, totalFailed *int, totalSkipped *int) TestSuiteReport {
	var testCases []TestCaseReport
	for _, testcase := range testsuite.Testcases {
		var result TestStatus
		if testcase.IsSkipped() {
			*totalSkipped++
			result = StatusSkipped
		} else if testcase.HasFailed() {
			*totalFailed++
			result = StatusFailed
		} else {
			result = StatusPassed
		}
		testCases = append(testCases, TestCaseReport{
			Name:   testcase.Name,
			Result: result,
		})
	}
	return TestSuiteReport{
		Name:          testsuite.Name,
		TestCases:     testCases,
	}
}

func WriteXMLToFile(report Report, filePath string) error {
	data, err := xml.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
