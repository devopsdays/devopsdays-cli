# Changelog - devopsdays-cli

### 0.15.1

__Changes__

- Update goreleaser to make it a prerelease

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- Update README with video

__Contributors__

- Matt Stratton

Released by Matt Stratton, Thu 05 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.15.0...0.15.1#diff)
______________

### 0.15.0

__Changes__

- Ensure that event content path exists

  Fixes #90

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>





__Contributors__

- Matt Stratton

Released by Matt Stratton, Thu 05 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.14.0...0.15.0#diff)
______________

### 0.14.0

__Changes__

- Move banner call

  This gives a little more flexability, and it also handles the issue where banner was hijacking the help flags.

  Fixes #75

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- Add proper fetching of talks

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- Add small functionality to helpers.GetTalks

  It only returns the filenames right now, but it’s a start

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- Fix regression with event name

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- MVP of create event function

  Introduced bug - name field in YAML has the city name capitalized (it is not getting this value from the struct)

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- Fix twitter field in event template

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- Add helpers.CopyFile function

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- Add nouns and verbs to README

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- Remove dummy file
- Adding dummy file to test auto-sign

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- Add configuration for stale bot

  Signed-off-by: Matt Stratton <matt.stratton@gmail.com>
- Fix spacing error in speaker template
- Update event model
- Move templates into template file
- Implement scaffold of GetTalks function

__Contributors__

- Matt Stratton

Released by Matt Stratton, Thu 05 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.13.0...0.14.0#diff)
______________

### 0.13.0

__Changes__

- Update README
- Merge pull request #74 from devopsdays/add-survey

  Switch to using survey for prompts
- Switch to using survey for prompts

  We now use the survey package, which is a lot prettier. It also handles the use of an editor for the bio, etc, which is much MUCH better.

  This commit also enables the city and year flags on the create speaker command properly.

  Fixes #72
- Add bio for create speaker

  Also added bio to speaker model.
  Fixes #65
- Remove old docs folder

__Contributors__

- Matt Stratton

Released by Matt Stratton, Wed 04 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.12.8...0.13.0#diff)
______________

### 0.12.8

__Changes__

- Various cobra cleanup and added banner
- Add speaker image creation feature
- WIP on the speaker image functionality
- Add coveralls
- Add more badges
- Fix badges

__Contributors__

- Matt Stratton

Released by Matt Stratton, Tue 03 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.12.7...0.12.8#diff)
______________

### 0.12.7

__Changes__

- Remove bintray for now

__Contributors__

- Matt Stratton

Released by Matt Stratton, Tue 03 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.12.6...0.12.7#diff)
______________

### 0.12.6

__Changes__

- Change to dist directory to see if travis can find it

__Contributors__

- Matt Stratton

Released by Matt Stratton, Tue 03 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.12.5...0.12.6#diff)
______________

### 0.12.5

__Changes__

- Fix i386 upload to bintray

__Contributors__

- Matt Stratton

Released by Matt Stratton, Tue 03 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.12.4...0.12.5#diff)
______________

### 0.12.4

__Changes__

- Whoops

__Contributors__

- Matt Stratton

Released by Matt Stratton, Tue 03 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.12.3...0.12.4#diff)
______________

### 0.12.3

__Changes__

- Troubleshoot travis

__Contributors__

- Matt Stratton

Released by Matt Stratton, Tue 03 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.12.2...0.12.3#diff)
______________

### 0.12.2

__Changes__

- Add rpm deploy

__Contributors__

- Matt Stratton

Released by Matt Stratton, Tue 03 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.12.1...0.12.2#diff)
______________

### 0.12.1

__Changes__

- Small updates for deploy

__Contributors__

- Matt Stratton

Released by Matt Stratton, Tue 03 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.12.0...0.12.1#diff)
______________

### 0.12.0

__Changes__

- Implement create speaker functionality

  Fixes #41
- Change most tests to convey tests
- Add some convey tests
- Scaffold out beginnings of new speaker code
- Remove unnecessary files
- Merge pull request #62 from devopsdays/data-model

  Migrate to data model and use helpers package
- Move all helpers to other package
- Initialize model package
- Add a couple more scaffolds to sponsor and talk

__Contributors__

- Matt Stratton

Released by Matt Stratton, Tue 03 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.11.2...0.12.0#diff)
______________

### 0.11.2

__Changes__

- Make version script use better commit tense

__Contributors__

- Matt Stratton

Released by Matt Stratton, Mon 02 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.11.1...0.11.2#diff)
______________

### 0.11.1

__Changes__

- Fix docs and help

__Contributors__

- Matt Stratton

Released by Matt Stratton, Mon 02 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.11.0...0.11.1#diff)
______________

### 0.11.0

__Changes__

- Scaffold out nouns and verbs

__Contributors__

- Matt Stratton

Released by Matt Stratton, Mon 02 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.20...0.11.0#diff)
______________

### 0.10.20

__Changes__

- Try fixing package name

__Contributors__

- Matt Stratton

Released by Matt Stratton, Mon 02 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.19...0.10.20#diff)
______________

### 0.10.19

__Changes__

- Fix travis vars

__Contributors__

- Matt Stratton

Released by Matt Stratton, Mon 02 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.18...0.10.19#diff)
______________

### 0.10.18

__Changes__

- Change to use direct API instead of CLI for bintray

__Contributors__

- Matt Stratton

Released by Matt Stratton, Mon 02 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.17...0.10.18#diff)
______________

### 0.10.17

__Changes__

- Move bintray back to make

__Contributors__

- Matt Stratton

Released by Matt Stratton, Mon 02 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.16...0.10.17#diff)
______________

### 0.10.16

__Changes__

- Fix up travis script

__Contributors__

- Matt Stratton

Released by Matt Stratton, Mon 02 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.15...0.10.16#diff)
______________

### 0.10.15

__Changes__

- Make travis great again

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sun 01 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.14...0.10.15#diff)
______________

### 0.10.14

__Changes__

- Make travis simpler

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sun 01 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.13...0.10.14#diff)
______________

### 0.10.13

__Changes__

- Dummy commit

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sun 01 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.12...0.10.13#diff)
______________

### 0.10.12

__Changes__

- Remove stray junk from travis

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sun 01 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.11...0.10.12#diff)
______________

### 0.10.11

__Changes__

- Stupid yaml

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sun 01 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.10...0.10.11#diff)
______________

### 0.10.10

__Changes__

- Stupid jfrog

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sun 01 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.9...0.10.10#diff)
______________

### 0.10.9

__Changes__

- embed jfrog bin

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sun 01 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.8...0.10.9#diff)
______________

### 0.10.8

__Changes__

- Add jfrom client again

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sun 01 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.7...0.10.8#diff)
______________

### 0.10.7

__Changes__

- Add bintray client install to travis

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sun 01 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.6...0.10.7#diff)
______________

### 0.10.6

__Changes__

- Add packages to bntray upload

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sun 01 Oct 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.5...0.10.6#diff)
______________

### 0.10.5

__Changes__

- Add fpm back to goreleaser

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.4...0.10.5#diff)
______________

### 0.10.4

__Changes__

- Switch back to container

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.3...0.10.4#diff)
______________

### 0.10.3

__Changes__

- Fix travis - doesn't need docker login

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.2...0.10.3#diff)
______________

### 0.10.2

__Changes__

- Add DODPATH env variable to travis

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.1...0.10.2#diff)
______________

### 0.10.1

__Changes__

- Attempt docker update for fpm in travis

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.10.0...0.10.1#diff)
______________

### 0.10.0

__Changes__

- Make better top level verbs

  Also moved the doctor command under show config.
  Fixes #34











































































__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.9.9...0.10.0#diff)
______________

### 0.9.9

__Changes__

- Update README with proper path to brew tap

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.9.8...0.9.9#diff)
______________

### 0.9.8

__Changes__

- Make README nicer

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.9.7...0.9.8#diff)
______________

### 0.9.7

__Changes__

- Remove fpm for now

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.9.6...0.9.7#diff)
______________

### 0.9.6

__Changes__

- Specify ruby in goreleaser step

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.9.5...0.9.6#diff)
______________

### 0.9.5

__Changes__

- More travis troubleshooting

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.9.4...0.9.5#diff)
______________

### 0.9.4

__Changes__

- Add ruby gems to path on travis

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.9.3...0.9.4#diff)
______________

### 0.9.3

__Changes__

- Troubleshoot fpm on travis

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.9.2...0.9.3#diff)
______________

### 0.9.2

__Changes__

- Add next step of version script

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.9.1...0.9.2#diff)
______________

### 0.9.1

__Changes__

- Update version script one step

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.9.0...0.9.1#diff)
______________

### 0.9.0

__Changes__

- Update version script

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.8.9...0.9.0#diff)
______________

### 0.8.9

__Changes__

- Add changelog
- Remove old changelog
- Remove test file
- Fix changelog stuff
- Extra file
- Fix extra space in gump script
- Add changelog and README automation

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.8.8...0.8.9#diff)
______________

### 0.8.8

__Changes__

- Add fpm to goreleaser again & move homebrew tap

__Contributors__

- Matt Stratton

Released by Matt Stratton, Sat 30 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.8.7...0.8.8#diff)
______________

### 0.8.7

__Changes__

- What the eff with travis

__Contributors__

- Matt Stratton

Released by Matt Stratton, Fri 29 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.8.6...0.8.7#diff)
______________

### 0.8.6

__Changes__

- fix travis
- Update readme

__Contributors__

- Matt Stratton

Released by Matt Stratton, Fri 29 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.8.5...0.8.6#diff)
______________

### 0.8.5

__Changes__

- Remove fpm stuff

__Contributors__

- Matt Stratton

Released by Matt Stratton, Fri 29 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.8.4...0.8.5#diff)
______________

### 0.8.4

__Changes__

- fix travis?

__Contributors__

- Matt Stratton

Released by Matt Stratton, Fri 29 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.8.3...0.8.4#diff)
______________

### 0.8.3

__Changes__

- Clean up travis

__Contributors__

- Matt Stratton

Released by Matt Stratton, Fri 29 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.8.2...0.8.3#diff)
______________

### 0.8.2

__Changes__

- Clean up travis

__Contributors__

- Matt Stratton

Released by Matt Stratton, Fri 29 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.8.1...0.8.2#diff)
______________

### 0.8.1

__Changes__

- Attempt packaging

__Contributors__

- Matt Stratton

Released by Matt Stratton, Fri 29 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/0.8.0...0.8.1#diff)
______________

### 0.8.0

__Changes__

- Build travis
- Add a bunch of crap
- Revert version script
- update version script for readme stuff
- Add doctor and version features and much more
- Add additional data to version command

  Fixes #23
- Update appveyor
- Remove changelog generator for now
- Combine appveyor and travis
- Merge branch 'refactor-cobra'
- Fix version script
- Update readme and some tools
- add templated README
- Update main.go with package description
- Fix appveyor for the millionth time
- move to after build
- clean up via spinner idea
- Try gh releases...AGAIN
- Fix gh release thing for appveyor
- remove travis deploy to test appveyor
- Remove github-releases step
- Try uploads
- add some quotes for fun
- Try quotes
- Troubleshoot github-release
- Try an old tagger
- Hardcode tag for release upload
- Try to fix upload again
- Add beginning of new cobra
- Use version for tag
- Add long form tag flag
- oh yeah, ponyville
- try using github-release in appveyor
- Try to force GH release upload
- Trim down appveyor
- Remove rdp and clone folder steps
- Add after build troubleshooting step
- Add rdp for troubleshooting
- More appveyor debugging
- Remove stupid build stupid stupid
- hardcode msi directory pasth
- Fix bad dir statement in appveyor command
- Try using a build directory
- Attempt another appveyor config
- Try makefile artifact
- Appveyor troubleshooting
- Add all msi to deploy
- Force artifact version
- find artifacts in appveyor
- Add artificats for appveyor
- Clean up appveyor build
- Add support for MSI to appveyor build
- Appvery fun
- Add user to appvey
- Change to use tokens in appveyor
- Update appveyor
- More fun with builds
- Remove test file
- Try with newer github release process
- Fix license file
- add go get to appveyor
- Add test file
- Remove matt
- Remove glide stuff
- Dump thing
- Update appveryer
- Update version script with proper name
- Add Appveyor stuff for Windows
- Make draft release prettier
- Add version script
- Add homebrew tap
- Update README with proper badge links
- Add version into build
- Fix mousetrap for real
- Fix mousetrap
- Try to fix version thing
- Try with goreleaser again
- Remove stray prompt junk
- Tweak some version stuff
- Revert use of prompt
- Remove windows
- Update travis
- Merge pull request #28 from devopsdays/add-prompt

  Start using prompt
- Start using prompt
- Fix makefile
- Update build scripts
- Add poc of prompt
- Try a release
- Rename to devopsdays-cli
- Fix makefile
- Add sample data for testing
- remove goveralls because travis failing
- Update travis and readme
- Change package name to devopsdays

  Fixes #26
- Update Appveyor release stuff
- First attempt at reflecting on TeamMembers
- add city and year flags
- Add function to list organizers
- Remove tag builds since appveyor was going nuts
- Add go get for deps in Appveyor
- Add appveyor release step
- Add bio to struct
- Implement edit event command

  MVP, but it mostly works
- First implementation of event edit
- Merge branch 'master' of github.com:mattstratton/probablyfine
- Put rice file in the proper location
- Put rice file in the proper location
- Clean up and update event data file
- Remove goveralls step for now
- Embed templates using rice tool
- Make travis release properly
- Update go version
- Better travis trigger
- Trigger travis verbose
- Fix tests
- Accomplish basic new event creation
- Change license to MIT
- Add doctor command to confirm hugo version
- Refactor eventDataPath to pass in webdir
- Add config command
- Add setWebdir function
- Update version command to remove Hugo
- Add version number to Makefile
- Merge branch 'master' of github.com:mattstratton/probablyfine
- Add version command
- Update README.md
- Clean up docs and add badge
- Merge branch 'master' of github.com:mattstratton/probablyfine
- Implement create event (data file only)
- Implement create event (data file only)
- Add some documentation
- Clean up some linting errors
- Merge pull request #8 from mattstratton/mattstratton/implement-cobra

  Change from cli to cobra
- Add mousetrap dep for release
- Change from cli to cobra
- Clean up some comments
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Add image resizing for sponsors and move templates
- Add command to add sponsor
- Add year tests
- Increment version
- try to fix travis again
- remove makefile
- revert travis
- fix travis again
- Update travis file
- remove releases stuff
- Merge branch 'master' of github.com:mattstratton/probablyfine

  # Conflicts:
  #       .travis.yml
- Making travis work
- Making travis work
- Add coveralls to travis
- Add godoc comments for addEvent()
- Add city, year, twitter validation
- Add travis badge
- Fix os in travis file
- Add base travis file
- Add badges
- Add logo to README
- Merge pull request #7 from mattstratton/mattstratton/golang-rewrite

  Switch from ruby to golang
- Add some documentation
- Add tests for validateField()
- Write TODO content for addSponsor()
- Add tests for cityClean and eventDataPath
- Added command to create event
- Add first scaffold of golang app and tests
- Remove all ruby files and prep for golang
- Add some more code for sponsors
- Worked on getting some POC going of the generating event stuff
- Add stuff to the gemspect because it beats adding functionality
- Create sample proposal migration script
- Scaffold out subcommands
- Add a bunch of help stuff
- Merge branch 'master' of github.com:mattstratton/probablyfine
- Initial commit and scaffold
- Initial commit

__Contributors__

- Matt Stratton

Released by Matt Stratton, Fri 29 Sep 2017 -
[see the diff](https://github.com/mh-cbon/devopsdays-cli/compare/1afb5ad272771069a0d86e66314e132bfcce377c...0.8.0#diff)
______________


