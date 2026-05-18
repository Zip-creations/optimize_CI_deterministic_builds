package main

import "fmt"
import cfg "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/config"
import junit "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/jUnit"
import disc "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/testDiscovery"
import out "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/xmlOutput"

func main() {
	// Read config
	config, err := cfg.ReadConfig("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("config:\n", config, "\n")  // Debug

	// Read all existing tests from the user-configured script
	allSuites, err := disc.RunTestDiscoveryScript(config.TestDiscoveryPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Suits from discovery:\n", allSuites, "\n")  // Debug

	// Read all tests in the JUnit XML output of the last run (if existing)
	allSuitesJUnit, err := junit.ReadJUnitTestSuites(config.JUnitXMLDirectory)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Suites from JUnit XML:\n", allSuitesJUnit, "\n")  // Debug

	report := out.MatchTests(allSuites, allSuitesJUnit)
	out.WriteXMLToFile(report, config.OutputPath)
	// fmt.Println("Successfully created report: \n", report)  // Debug
}
