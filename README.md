# Daiv Worklog

A plugin for the daiv CLI tool.

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

- worklog.directory: The directory where the worklog files are stored.

You can configure these settings when you first run daiv after installing the plugin.

## Usage

After installation, the plugin will be automatically loaded when you start daiv.

