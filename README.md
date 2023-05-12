# USVisaWaitTimes
A tool for checking US visa wait times.

## Usage
### Install US-Visa with Homebrew (macOS Only)
```bash
brew tap owo-network/brew
brew install us-visa
```
### Check the embassies where you can apply for a U.S. visa.
```bash
./us-visa -l shanghai
Shanghai
```
### Use fuzzy query embassies
```bash
./us-visa -l sh
Ashgabat
Bishkek
Dushanbe
Kinshasa
Shanghai
Shenyang
Tashkent
```
### Check the schedule of the designated embassy
```bash
./us-visa -s shanghai
Interview Required Visitors (B1/B2): 141 Calendar Days
Interview Required Students/Exchange Visitors (F, M, J): 30 Calendar Days
Interview Required Petition-Based Temporary Workers (H, L, O, P, Q): 35 Calendar Days
Interview Required Crew and Transit (C, D, C1/D): 36 Calendar Days
Interview Waiver Students/Exchange Visitors (F, M, J): 1 Calendar Days
Interview Waiver Petition-Based Temporary Workers (H, L, O, P, Q): 1 Calendar Days
Interview Waiver Crew and Transit (C, D, C1/D): 1 Calendar Days
Interview Waiver Visitors (B1/B2): 1 Calendar Days

Made with ♥ by Vincent.
Data from https://travel.state.gov
GitHub: https://github.com/missuo/USVisaWaitTimes
```

## Author
**USVisaWaitTimes** © [Vincent Young](https://github.com/missuo), Released under the [MIT](./LICENSE) License.<br>