# Advent of Code 2022

<https://adventofcode.com/>

## Solutions

|  Day  | Solution 1 | Solution 2 | Duration |
| ----- | ---------- | ---------- | -------- |
{{- range $day, $solutions := .results }}
| {{ $day }} | {{ $solutions.First }} | {{ $solutions.Second }} |  {{ $solutions.Duration }} |
{{- end }}
| ----- | ---------- | ---------- | -------- |
|       |            |Total Time: | {{ .totalDuration }} |
