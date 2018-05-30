package compiler

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file3 := &embedded.EmbeddedFile{
		Filename:    "babel-standalone-bower/.git",
		FileModTime: time.Unix(1520881882, 0),
		Content:     string("gitdir: ../../../../.git/modules/js2/compiler/lib/babel-standalone-bower\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "babel-standalone-bower/README.md",
		FileModTime: time.Unix(1520881903, 0),
		Content:     string("You probably want to go to https://github.com/Daniel15/babel-standalone instead.\n\nThis repo is only required because Bower requires a Git repo containing compiled code, which is silly. Build artifacts don't belong in Git."),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "babel-standalone-bower/babel.js",
		FileModTime: time.Unix(1520881903, 0),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "babel-standalone-bower/babel.min.js",
		FileModTime: time.Unix(1520881903, 0),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "babel-standalone-bower/bower.json",
		FileModTime: time.Unix(1520881903, 0),
		Content:     string("{\n  \"name\": \"babel-standalone\",\n  \"homepage\": \"https://github.com/Daniel15/babel-standalone\",\n  \"authors\": [\n    \"Daniel Lo Nigro <daniel@dan.cx>\"\n  ],\n  \"description\": \"Standalone build of Babel for use in non-Node.js environments, including browsers. Similar to the (now deprecated) babel-browser.\",\n  \"main\": \"babel.js\",\n  \"keywords\": [\n    \"babel\",\n    \"babeljs\",\n    \"6to5\",\n    \"transpile\",\n    \"transpiler\"\n  ],\n  \"license\": \"MIT\"\n}\n"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    "babel-standalone-bower/update.ps1",
		FileModTime: time.Unix(1520881903, 0),
		Content:     string("\ufeff# Checks if newer versions of Babel are available on NPM, and updates this repo if so.\n\nSet-StrictMode -Version Latest\n\n# Checks that the last ran command returned with an exit code of 0\nfunction Assert-LastExitCode {\n  if ($LASTEXITCODE -ne 0) {\n    throw 'Non-zero exit code encountered'\n  }\n}\n\n# Converts a Windows path (\"C:\\foo\") to a Cygwin path (\"/cygdrive/c/foo\") for usage with \"tar\" from\n# Gnuwin32\nfunction Get-CygwinPath([Parameter(Mandatory)][String] $Path) {\n  '/cygdrive/' + $Path.Replace(':', '').Replace('\\', '/')\n}\n\n# Expands a .tar.gz archive\nfunction Expand-GzipArchive(\n  [Parameter(Mandatory)][String] $Path,\n  [Parameter(Mandatory)][String] $DestinationPath\n) {\n  if ([Environment]::OSVersion.Platform -eq [PlatformID]::Win32NT) {\n    # On Windows, we're using a Cygwin version of tar, so the paths need to be adjusted.\n    $Path = Get-CygwinPath $Path\n    $DestinationPath = Get-CygwinPath $DestinationPath\n  }\n  tar zvxf $Path -C $DestinationPath; Assert-LastExitCode\n}\n\n# Checks if the specified Git tag exists\nfunction Test-GitTag([Parameter(Mandatory)] [String] $Tag) {\n  git rev-parse $Tag > $null 2>&1\n  $LASTEXITCODE -eq 0\n}\n\ngit pull; Assert-LastExitCode\n\n$npm_data = Invoke-RestMethod -Uri https://registry.npmjs.org/babel-standalone\n$new_versions = $npm_data.versions.PSObject.Members | \n  # Get properties of \"versions\" map where a Git tag does not already exist\n  Where-Object { $_.MemberType -eq 'NoteProperty' -and -Not (Test-GitTag -Tag ('v' + $_.Name))} |\n  ForEach-Object {\n    [PSCustomObject]@{\n      DownloadUrl = $_.Value.dist.tarball\n      Version = ([Version]$_.Name)\n    }\n  } |\n  Sort-Object -Property Version\n\nforeach ($version in $new_versions) {\n  Write-Output (\"{0}...\" -f $version.Version)\n  # Download archive from NPM\n  $temp_file = New-TemporaryFile\n  Invoke-WebRequest -Uri $version.DownloadUrl -OutFile $temp_file\n  \n  # Extract to temporary directory\n  $temp_dir = [IO.Path]::GetTempPath() + [IO.Path]::GetRandomFileName()\n  New-Item -ItemType Directory -Path $temp_dir > $null\n  Expand-GzipArchive -Path $temp_file -DestinationPath $temp_dir\n  \n  # Grab just the files we care about\n  Copy-Item -Path ([IO.Path]::Combine($temp_dir, 'package', 'babel.js')) -Destination .\n  Copy-Item -Path ([IO.Path]::Combine($temp_dir, 'package', 'babel.min.js')) -Destination .\n  Remove-Item -Path $temp_file\n  Remove-Item -Path $temp_dir -Recurse\n  \n  # Update Git\n  git commit -m ('Upgrade to Babel {0}' -f $version.Version) --author='DanBuild <build@dan.cx>' babel.js babel.min.js; Assert-LastExitCode\n  git tag -a ('v' + $version.Version) -m ('Automated upgrade to Babel {0}' -f $version.Version); Assert-LastExitCode\n}\n\ngit push origin master --follow-tags; Assert-LastExitCode"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1520881879, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dir2 := &embedded.EmbeddedDir{
		Filename:   "babel-standalone-bower",
		DirModTime: time.Unix(1520881903, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file3, // "babel-standalone-bower/.git"
			file4, // "babel-standalone-bower/README.md"
			file5, // "babel-standalone-bower/babel.js"
			file6, // "babel-standalone-bower/babel.min.js"
			file7, // "babel-standalone-bower/bower.json"
			file8, // "babel-standalone-bower/update.ps1"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir2, // "babel-standalone-bower"

	}
	dir2.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`lib`, &embedded.EmbeddedBox{
		Name: `lib`,
		Time: time.Unix(1520881879, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
			"babel-standalone-bower": dir2,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"babel-standalone-bower/.git":         file3,
			"babel-standalone-bower/README.md":    file4,
			"babel-standalone-bower/babel.js":     file5,
			"babel-standalone-bower/babel.min.js": file6,
			"babel-standalone-bower/bower.json":   file7,
			"babel-standalone-bower/update.ps1":   file8,
		},
	})
}