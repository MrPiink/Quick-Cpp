# Task List for Release Process Improvement

## Remove Extra Tar Files

- Remove all unnecessary tar files from the assets, keeping only the Windows tar.

## Automate Pre-Releases via CLI

- Configure the CLI to always create a pre-release when initiating a release.

## Implement Versioning Scheme

- Start from an initial version, e.g., `0.1.0`.
- Use release candidates (e.g., `0.1.1-rc1`) for testing releases.
- Allow for additional commits without affecting the final version number if a release candidate fails.
- Upon success, update the version to the next appropriate number (e.g., `0.1.2`).

## Continuous Integration (CI) Enhancements

- Set up the CI to create new tags (e.g., `0.1.2`) upon successful builds.
- Push these tags as official releases.

## Manual Release Management

- Decide whether to manually add release information or automate it through the CI.
- Initially set releases as pre-releases.
- After reviewing and adding necessary changes, manually switch the pre-release to a full release if appropriate.

## Manage Release Candidate Tags

- Keep release candidate (`rc`) tags for testing purposes.
- Clean up obsolete `rc` tags periodically to avoid clutter.
- Ensure only meaningful official releases are displayed prominently.

## Executable Packaging

- Place the executable into the `tools` folder before cleaning up unnecessary files.
- Verify that the executable remains in the `tools` folder after file removal.
- Configure the pipeline to package the application using the contents of the `tools` folder.

## Automated Deployment Setup

- Adjust the pipeline to facilitate automated deployment.
- Hold off on publishing to the community feed until the deployment process is fully tested and stable.
