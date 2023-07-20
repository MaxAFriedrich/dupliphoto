# Dupliphoto

A simple hash based photo collation and merge program for UNIX systems.

This program organizes your photos in deep folder structures with random file names. It requires two lists: one for target folders and another for their corresponding source folders. The program will find images in the source folders and move them to the target folders with a consistent naming scheme. It simplifies managing large collections of photos.

## Usage

Create a config file in the form:

```yaml
blocks:
  - target: /path/to/target_folder
    sources:
      - /path/to/source_folder1
      - /path/to/source_folder2
      - /path/to/source_folder3

  - target: /path/to/another_target_folder
    sources:
      - /path/to/another_source_folder1
      - /path/to/another_source_folder2
      - /path/to/another_source_folder3
```

Then run:

``` bash
dupliphoto config.yml
```

For additional options run:

```bash
dupliphoto --help
```


