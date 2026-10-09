# Radforge demo

A small Go repository for trying Radforge cloud analysis. All data and code are synthetic. There are no external dependencies or credentials.

The shipping policy is free shipping for orders of at least 5,000 cents. The tests check small and large orders but deliberately omit the threshold and values around it. Passing tests therefore do not establish that the policy is protected.

Run the example locally:

```sh
go test ./...
```

The [demo pull request](https://github.com/Radforge-Dev/radforge-demo/pull/1) deliberately violates that policy. Its **Ordinary Go tests** check runs the tests above. This check is separate from Radforge analysis; a green result only means those two example cases pass.

## Try the GitHub App

1. Create your account at [Radforge](https://radforge.dev) and select a repository you administer. To experiment independently, use a fork of this repository.
2. Complete GitHub linking and confirm the repository is enabled in Radforge Setup.
3. Open a pull request changing `shipping.go`, such as raising the threshold from 5,000 to 7,500 cents. The current tests still pass, although orders between those values now violate the documented policy.
4. Read the Radforge check for that pull request. Follow its findings link and verify the analyzed commit.
5. Add a test for the original threshold and nearby values, correct the code, and push another commit. Each admitted analysis uses your shared trial allowance.

See the [GitHub first-run guide](https://docs.radforge.dev/github-app/installation/) for supported setup and trial limits. A queued check is not a completed analysis. No cloud-analysis outcome is recorded in this example yet.

## Deliberate test gaps

The table omits 4,999, 5,000, and 5,001 cents. A change from `>=` to `>` can pass the existing tests. The demo pull request also changes the threshold while leaving the intended policy unchanged.

This is a teaching example, not production billing or shipping logic.
