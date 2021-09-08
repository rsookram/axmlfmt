# axmlfmt

axmlfmt is an opinionated formatter for Android XML resources. It takes XML
that looks like

```xml
<?xml version="1.0" encoding="utf-8"?>
<LinearLayout
    xmlns:tools="http://schemas.android.com/tools"
    android:layout_height="match_parent"
    android:layout_width="match_parent"
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:orientation="vertical"
    tools:context=".MainActivity">


    <!-- Not using AppCompat to reduce app size -->
  <Toolbar
      android:layout_width="match_parent"
      android:id="@+id/toolbar"
      android:layout_height="wrap_content"
      android:title="@string/app_name"/>

    <ScrollView
        android:layout_weight="1"
        android:layout_width="match_parent"
      android:layout_height="0dp"
        android:fillViewport="true">

        <EditText
                android:id="@+id/text"
            android:layout_width="match_parent"
            android:importantForAutofill="no"
            android:layout_height="wrap_content"
            android:textSize="16sp">
          </EditText>

    </ScrollView>
</LinearLayout>
```

and turns it into

```xml
<?xml version="1.0" encoding="utf-8"?>
<LinearLayout
    xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:tools="http://schemas.android.com/tools"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:orientation="vertical"
    tools:context=".MainActivity">

    <!-- Not using AppCompat to reduce app size -->
    <Toolbar
        android:id="@+id/toolbar"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:title="@string/app_name" />

    <ScrollView
        android:layout_width="match_parent"
        android:layout_height="0dp"
        android:fillViewport="true"
        android:layout_weight="1">

        <EditText
            android:id="@+id/text"
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:importantForAutofill="no"
            android:textSize="16sp" />
    </ScrollView>
</LinearLayout>
```


## Install

Pre-compiled binaries of axmlfmt can be downloaded from
[the releases page](https://github.com/rsookram/axmlfmt/releases).
Alternatively, you can install it with with the
[`go` command](https://golang.org/doc/install) by running:

```shell
go install github.com/rsookram/axmlfmt/cmd/axmlfmt@latest
```


## Usage

One way you may want to use axmlfmt is to have it format all the XML files in a
git repository. This can be done with the following command:

```shell
git ls-files '*.xml' | xargs axmlfmt -w
```

The full usage description is:

```
USAGE:
    axmlfmt [FLAGS] [FILE]...

FLAGS:
    -h, -help, --help    Prints help information
    -V                   Prints version information
    -w                   Writes result to (source) file instead of stdout

ARGS:
    <FILE>...    Path of XML files to format
```


## Build

axmlfmt can be built from source by cloning this repository and using the `go`
command.

```shell
git clone https://github.com/rsookram/axmlfmt
cd axmlfmt
go build ./cmd/axmlfmt
```


License
-------

    Copyright 2019 Rashad Sookram

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
