# Branch protection setup

Apply these settings on GitHub after the first CI workflow runs on `main`.

## Required settings

- Default branch: `main`
- Require pull request before merging
- Require status checks to pass:
  - `CI / Go 1.22`
  - `CI / Go stable`
  - `CI / govulncheck`
- Require branches to be up to date before merging
- Do not allow force pushes
- Do not allow deletions

## First push to `main`

```bash
git checkout main
git remote add origin https://github.com/VZaps/vzaps-sdk-go.git
git push -u origin main
```

Create the remote repository first if it does not exist.

## Release workflow

Tag a semver release to publish on GitHub and index on pkg.go.dev:

```bash
git tag v0.1.0
git push origin v0.1.0
```

## Automated setup (GitHub CLI)

Install [GitHub CLI](https://cli.github.com/) and run from the repository root:

```bash
gh api repos/{owner}/{repo}/branches/main/protection \
  --method PUT \
  --field required_status_checks[strict]=true \
  --field required_status_checks[contexts][]='CI / Go 1.22' \
  --field required_status_checks[contexts][]='CI / Go stable' \
  --field required_status_checks[contexts][]='CI / govulncheck' \
  --field enforce_admins=true \
  --field required_pull_request_reviews[required_approving_review_count]=1 \
  --field restrictions=null
```

Adjust `contexts` to match the exact job names shown in the Actions tab after the first green run.
