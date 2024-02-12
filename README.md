## code-hc

This is a simple command line code health checker using LLM and git commands. 

### TODOs:
- [ ] Handle maximum token length
- [ ] Enable change of model
- [ ] Enable storage and retrieval of context and chat session
- [ ] Enable image input with a simple interface
- [ ] Add more flags to support different code review scenarios
- [ ] Better bot message parsing and formatting

### demo
```sh
code-hc -h # or run with "go run main.go -h"
```
```md
Usage of code-hc:
  -g string
        Git show/diff/archive command
        Example: -g "git show" or -g "git diff" or -g "git archive HEAD [path]"
```

```sh
code-hc
```
```md
Git command not found. Default to git diff
Output from [git diff] command: 55 lines
Enter instruction for the bot:                           

**Bot**: From the provided diff, the code appears to be well-structured and readable. However, there are a few points that could be improvement areas for better code health:

1. Consistent Error Messages: ...

2. Error Handling: ...

3. Using Default Git Command: ...

4. Avoid Terminating the Program in `init()`: It's generally not good practice to call `os.Exit(1)` inside the `init()` function. ...

5. Also, the typo in the `fmt.Println("Git command not fonud. Default to git show")` can be corrected to `"Git command not found. Default to git show"`.

...

**User**: why should I avoid terminating the program in init()?

**Bot**: There are a couple of reasons why you might want to avoid using `os.Exit` in your `init()` function.

1. Testability: ...

2. Cleanliness: ...

3. Debugging: ...

4. Order of execution: ...

A better approach would be to ...

**User**: q (quit)
```


```sh
code-hc -g "git show dc5710d8d9ec351aa1b2c4ef3e2022cc094797f6"
```
```md
Output from [git show dc5710d8d9ec351aa1b2c4ef3e2022cc094797f6] command: 160 lines
Enter instruction for the bot: what did I do in this commit?

**Bot**: The commit made a significant rework on the structure of your agent code to improve the interaction between user and AI. Let's summarize the changes and provide some feedback:

...

However, some improvements and simplifications could be done:

...

**User**: what new features would you build on top of this MVP?

**Bot**: There are several features that could be built on top of this MVP to further enhance its capabilities and user experience:

...
```