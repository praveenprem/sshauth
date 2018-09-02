# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.1.0] - 02-09-2018
### Added
 - Unknown login attempt alert.
 - Initial SQL package (none-functional).
 - Changelog file.
 
### Changed
 - Alert package refactored to isolated service based package.
 - `SlackPayloadPetty` struct refactored.
 - README.md updated.
 
### Removed
 - HipChat function from Slack package.

## [1.0.0] - 13-07-1018
### Added
 - GitHub (Organisation / Team) based authentication.
 - GitLab Self-Hosted (Organisation / Group) based authentication.
 - Slack alerting support for authenticated with plugin.
 - System logging support for debugging.
 - Configuration and Exit Code Dictionary added for support and debugging.
 - SQL template.
  