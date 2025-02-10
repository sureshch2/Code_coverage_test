package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "sort"
)

// FileCoverage represents a single file's test coverage details
type FileCoverage struct {
    Name   string `json:"name"`
    Totals struct {
        Lines   int     `json:"lines"`
        Hits    int     `json:"hits"`
        Misses  int     `json:"misses"`
        Coverage float64 `json:"coverage"`
    } `json:"totals"`
}

// CodecovResponse represents the full response from Codecov API
type CodecovResponse struct {
    Totals struct {
        Coverage float64 `json:"coverage"`
    } `json:"totals"`
    Files []FileCoverage `json:"files"`
}

func main() {
    apiToken := "Your_codecov.io_token"  // Replace with your actual Codecov API token
    url := "https://api.codecov.io/api/v2/gh/openshift/repos/backplane-cli/report"

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("❌ Error creating request:", err)
        return
    }

    req.Header.Set("Authorization", "Bearer "+apiToken)
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("❌ Error making request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("❌ Error reading response:", err)
        return
    }

    var data CodecovResponse
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Println("❌ Error parsing JSON:", err)
        return
    }

    // Sort files by ascending test coverage
    sort.Slice(data.Files, func(i, j int) bool {
        return data.Files[i].Totals.Coverage < data.Files[j].Totals.Coverage
    })

    // Generate report
    report := fmt.Sprintf("# Backplane-CLI Detailed Code Coverage Report\n\n")
    report += fmt.Sprintf("## 📊 Overall Coverage\n- **Total Coverage**: `%.2f%%`\n\n", data.Totals.Coverage)
    report += "## 📉 Files with Low Test Coverage (Sorted in Ascending Order)\n"
    report += "| File | Total Lines | Covered Lines | Missed Lines | Coverage % |\n"
    report += "|------|------------|--------------|-------------|------------|\n"

    for _, file := range data.Files {
        report += fmt.Sprintf("| `%s` | %d | %d | %d | **%.2f%%** |\n",
            file.Name, file.Totals.Lines, file.Totals.Hits, file.Totals.Misses, file.Totals.Coverage)
    }

    err = ioutil.WriteFile("detailed_backplane_cli_coverage_report.md", []byte(report), 0644)
    if err != nil {
        fmt.Println("❌ Error writing report:", err)
        return
    }

    fmt.Println("✅ Test coverage report generated: detailed_backplane_cli_coverage_report.md")
}

