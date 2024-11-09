# Filestigator 🕵

## Short Summary
Filestigator is a CLI tool to help investigate a potentially malicious file by assigning a "suspicion score" based on findings like suspicious file name, malformed file extension, file content or magic numbers mismatch, and embedded script.

## Usage
TODO: Still in development, this will be written later

## Features
### File Name and Extension Analysis (In Progress)
Detect suspicious characteristics in file names, such as suspicious extensions (e.g. filename.php.jpg), null bytes and other dangerous characters.
### Magic Numbers and Content Inspection (In Progress)
Identify and detect mismatches between the file type declared in its magic numbers and file extension, as well as the actual content of the page.
### Embedded Script Detection (In Progress)
Detect a malicious script within a file that doesn't match its intended use (e.g. PHP code within an image file).
### Suspicion Ranking (In Progress)
Based on findings, the file will be scored and ranked. A file may be `SAFE`, `DOUBLE-CHECK`, `SUSPICIOUS`, and `DANGEROUS`.

## Docs
TODO: Might wanna write this later

## License
<a href="./LICENSE">LICENSE</a>