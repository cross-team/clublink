# Contributing
Everyone is welcome to join our community! When contributing to this repository, please first discuss the change you wish
to make via [Slack channel](https://short-d.com/r/slack) with the core team before making a change.

You can contribute to the project at your own pace & schedule. Please open a draft pull request after you started working 
on an issue so that everyone knows it's in progress. Please let the the core team know if you become too busy with life or lost 
interests in the project so that we can reassign the issue to other developers interested in making progress.

## Joining the Core Team
If you are dedicated to the project and want to have more impacts, you should consider joining the core team!
Please post your weekly contributions on our Slack every Sunday during your convenient hours. All of the followings counts as contributions: 

- Exploring the codebase
- Learnings new tools, infrastructures, libraries, algorithms, designs, or concepts
- Experimenting with different technologies and sharing the findings with the group
- Prototyping new features
- Fixing bugs
- Improving maintainability, scalability, reliability and performance of existing system
- Doing code reviews
- Writing design docs
- Creating UI mockups
- Organizing or participating sync meetings

After 2 months of consistent high quality contributions, you will receive an invitation to join the core team, 
where people together plan the product roadmap for Short, write design docs, build complex systems, create libraries, invent 
new distributed cloud infrastructure & drive innovations.

We are specificly looking for the following characters for the potential member:

- You are excited about learning new things and learning them very fast.
- You write clean code & constantly look for corners to simplify the existing systems.
- You believe in automation instead of doing everything manually in the long run.
- You strive for excellence in term of code quality, system design, and process optimization.
- You are a team player, who schedule pair programming sessions & solve difficult problems together with team members.
- You thrive in complexities & difficulties. Compilers, operating systems, programming languages, bitwise optimizations, 
distributed systems, machine learning vector optimization, computer vision algorithms excite you.
- You make things happen independently. Not dreaming or talking about them. You deliver them all the way onto production!

At the core team, the tech lead reserves the rights to move you back to general community after missing 3 consecutive standups.

## Pull Request Process

1. Update the [README](README.md) with details of changes to the interfaces,
   new environment variables, exposed ports, useful file locations and container
   parameters.
1. You may merge the Pull Request in once you have the sign-off of code owner,
   or if you do not have permission to do that, you may request the code owner
   to merge it for you.

## Code Review Guideline

### What to look for in a code review

- The code is well-designed.
- The functionality is good for the users of the code.
- Any UI changes are sensible and look good.
- Any parallel programming is done safely.
- The code isn’t more complex than it needs to be.
- The developer isn’t implementing things they might need in the future but don’t
  know they need now.
- Code has appropriate unit tests.
- Tests are well-designed.
- The developer used clear names for everything.
- Comments are clear and useful, and mostly explain why instead of what.
- Code is appropriately documented.

### The Standard

- Reviewers should favor approving a PR once it is in a state where it definitely
  improves the overall code health of the system being worked on, even if the PR
  isn’t perfect.
- Instead of seeking perfection, what a reviewer should seek is continuous
  improvement.
- If a PR adds a feature that the reviewer doesn’t want in their system, then the
  reviewer can certainly deny approval even if the code is well-designed.
- Reviewers should always feel free to leave comments expressing that something
  could be better, but if it’s not very important, prefix it with something like
  “Nit: “ to let the author know that it’s just a point of polish that they could
  choose to ignore.
- Checking in PRs that definitely worsen the overall code health of the system is
  not justified unless there is an emergency.

### Mentoring

Code review can have an important function of teaching developers something new
about a language, a framework, or general software design principles.

If the comment is purely educational, but not critical to meeting the standards
above, prefix it with "Nit: ".

### Principles

- Technical facts and data overrule opinions and personal preferences.
- On matters of style, the linters are the absolute authorities. The style should
  be consistent with what is there. If there is no previous style, accept the
  author’s.
- Aspects of software design are NOT personal preferences. Sometimes there are a
  few valid options. If the author can demonstrate that several approaches are
  equally valid, then the reviewer should accept the preference of the author.
  Otherwise the choice is dictated by standard principles of software design.
- If no other rule applies, then the reviewer may ask the author to be consistent
  with what is in the current codebase, as long as that doesn’t worsen the overall
  code health of the system.

### Showing Apprecaition

If you see something nice in the PR, tell the developer, especially when they
addressed one of your comments in a great way. Code reviews should offer
encouragement and appreciation for good practices, as well. It’s sometimes even
more valuable, in terms of mentoring, to tell a developer what they did right
than to tell them what they did wrong.

### Resolving Conflicts

When coming to consensus becomes especially difficult, it can help to have a
face-to-face meeting or a video call between the reviewer and the author, instead
of just trying to resolve the conflict through code review comments. (If you do
this, though, make sure to record the results of the discussion in a comment on
the CL, for future readers.)

Don’t let a PR sit around because the author and the reviewer can’t come to an
agreement.

Note: This guideline is derived from [Google Engineering Practices Documentation](https://github.com/google/eng-practices)

## Code of Conduct

- Using welcoming and inclusive language
- Being respectful of differing viewpoints and experiences
- Gracefully accepting constructive criticism
- Focusing on what is best for the community
- Showing empathy towards other community members

## Discussions

Please join this [Slack channel](https://short-d.com/r/slack) to
discuss bugs, dev environment setup, tooling, and coding best practices.
