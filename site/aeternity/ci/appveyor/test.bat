@echo on
@rem Matrix-driven Appveyor CI test script
@rem Currently only supports MSYS2 builds.
@rem See https://www.appveyor.com/docs/installed-software#mingw-msys-cygwin
@rem Required vars:
@rem    ERTS_VERSION
@rem    PACKAGE_TESTS_DIR
@rem    PACKAGE_ZIPARCHIVE
@rem    PLATFORM
@rem    TEST_STEPS
@rem    WIN_MSYS2_ROOT

SETLOCAL ENABLEEXTENSIONS
IF "%APPVEYOR_BUILD_FOLDER%"=="" (
	cd %~dp0\..\..\
) ELSE (
	cd %APPVEYOR_BUILD_FOLDER%
)
SET /p PACKAGE_VERSION=<%~dp0%\..\..\REVISION

rem Set required vars defaults

IF "%ERTS_VERSION%"=="" SET "ERTS_VERSION=9.3"
IF "%PACKAGE_TESTS_DIR%"=="" SET "PACKAGE_TESTS_DIR=/tmp/package_tests"
IF "%PACKAGE_ZIPARCHIVE%"=="" SET "PACKAGE_ZIPARCHIVE=aeternity-%PACKAGE_VERSION%-windows-x86_64.zip"
IF "%WIN_MSYS2_ROOT%"=="" FOR /F %%F IN ('where msys2') DO SET "WIN_MSYS2_ROOT=%%~dpF"
IF "%PLATFORM%"=="" SET "PLATFORM=x64"
IF "%TEST_STEPS%"=="" SET "TEST_STEPS=ct"

@echo Current time: %time%

rem SET the appropriate MSVC_VERSION
IF NOT "%MSVC_VERSION%"=="" GOTO MSVC_VERSION_SET
@FOR /F "tokens=* USEBACKQ delims=" %%F IN (`where /r "C:\Program Files (x86)\Microsoft Visual Studio" Microsoft.VCToolsVersion.default.txt`) DO SET "ToolsVerFile=%%F"
@FOR /F "tokens=* USEBACKQ delims=" %%F IN (`type "%ToolsVerFile%"`) DO SET MSVC_VERSION=%%F
:MSVC_VERSION_SET

rem Find and execute the VS env preparation script
@FOR /F "tokens=* USEBACKQ delims=" %%F IN (`where /r "C:\Program Files (x86)\Microsoft Visual Studio" vcvarsall`) DO SET vcvarsall="%%F"
call %vcvarsall% %PLATFORM%

rem Set the paths appropriately
SET PATH=%WIN_MSYS2_ROOT%\mingw64\bin;%WIN_MSYS2_ROOT%\usr\bin;%PATH%

:FIND_NEXT_STEP
set NEXT_STEP=
@for /f "tokens=1* delims=, " %%i in ("%TEST_STEPS%") do @(
	set "NEXT_STEP=%%i"
	set "TEST_STEPS=%%j"
)

IF "%NEXT_STEP%"=="" GOTO FINISHED_TESTING
GOTO TEST_%NEXT_STEP%

:TEST_ct
@echo Current time: %time%
rem Run test: ct
bash -lc "cd %BUILD_PATH% && epmd -daemon && make ct"

GOTO FIND_NEXT_STEP

:TEST_eunit
@echo Current time: %time%
rem Run test: eunit
bash -lc "cd %BUILD_PATH% && epmd -daemon && make eunit"

GOTO FIND_NEXT_STEP

:TEST_release
@echo Current time: %time%
rem Run test: release
bash -lc ^"cd %BUILD_PATH% ^&^& ^
           epmd -daemon ^&^& ^
           make python-env PIP=/mingw64/bin/pip3 ^&^& ^
           mkdir %PACKAGE_TESTS_DIR% ^&^& ^
           make python-release-test WORKDIR=%PACKAGE_TESTS_DIR% PACKAGE=`pwd`/%PACKAGE_ZIPARCHIVE% PYTHON=/mingw64/bin/python3^"

:FINISHED_TESTING

@echo Current time: %time%
rem Finished test phase
