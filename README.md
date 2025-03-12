# Daiv Worklog

A plugin for the [daiv](https://github.com/iures/daiv) CLI tool that helps you track and manage your work activities. This plugin provides seamless integration with daiv's standup functionality, allowing you to log your work and generate standup reports.

## Features
Reads all the files from the configured worklog directory, if they were creted or modified in the specified time range, it will add the contents of the file to the standup context.

## Installation

### From GitHub

```
daiv plugin install iures/daiv-worklog
```

### From Source

1. Clone the repository:
   ```
   git clone https://github.com/iures/daiv-worklog.git
   cd daiv-worklog
   ```

2. Build the plugin:
   ```
   go build -o out/daiv-worklog.so -buildmode=plugin
   ```

3. Install the plugin:
   ```
   daiv plugin install ./out/daiv-worklog.so
   ```

## Configuration

This plugin requires the following configuration:

- `worklog.directory`: The directory where the worklog files are stored.

You can configure these settings when you first run daiv after installing the plugin.

## Usage

After installation, the plugin will be automatically loaded when you run `daiv standup`.

### Standup Context

The plugin integrates with daiv's standup functionality, allowing you to generate reports of your work items within specific time ranges.
In order to do this, it looks for any files created in the configured worklog directory and uses the file modification time to determine when the work was logged.
The plugin then reads the contents of the file and adds to the standup context.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

