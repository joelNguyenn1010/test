## Puzzle Game

Following the Puzzle #1 example, write the execution flow based on the puzzle given.

#### Puzzle #1

```js
doSomething()
  .then(function() {
    return doSomethingElse();
  })
  .then(finalHandler);
```

Answer:

```
doSomething
|-----------------|
                  doSomethingElse(undefined)
                  |------------------|
                                     finalHandler(resultOfDoSomethingElse)
                                     |------------------|
```

#### Puzzle #2

```js
doSomething()
  .then(function() {
    doSomethingElse();
  })
  .then(finalHandler);
```

#### Puzzle #3

```js
doSomething()
  .then(doSomethingElse())
  .then(finalHandler);
```

#### Puzzle #4

```js
doSomething()
  .then(doSomethingElse)
  .then(finalHandler);
```

## Quick challenges

### Shell/CLI
1. What is happening in these Linux commands?  Describe as much as you know about what each symbol means and the effect it will have in the execution of the command, and the named programs.

   a) `(for i in {1..100}; do echo $i; done;) | grep 3 | grep -v 1 | paste -s -d+ - | bc`
   
   b) `[ ! -f /var/lock/myscript.lock ] && touch /var/lock/myscript.lock && (yum -y update >> /var/log/mylog.log 2>&1; ) && rm -f /var/lock/myscript.lock`


2. There is a directory, containing a large tree of subdirectories and files.  Scattered throughout these files are Australian phone numbers, and we want to harvest them – we want to end up with a simple list of the phone numbers.

   a) Write a regular expression to match Australian phone numbers.  The numbers will be in a mixture of the forms 02xxxxxxxx and +612xxxxxxxx, and there will also be the common usage of hyphens, spaces, and parentheses, so all of those common possibilities must be supported. eg. 02 xxxx xxxx, (02) xxxx-xxxx, +61 2 xxxx xxxx, +61 02 xxxxxxxx, +61 (0)2 xxxx-xxxx

   b) Write an example phone number, for each specific phone number format that your regex would match.

   c) Imagine that the full path of the directory is
   > /var/www/site1/uploads/phnumbers/
   
   Write a single-line or simple command that you could run from the shell (ideally Linux), to apply this regular expression to the files in the directory tree, and result in the simple list of phone numbers, one phone number per line.


### Software development

1.	Write a function/method/subroutine to determine if a string starts with an upper-case letter A-Z, *without* using any sort of isUpper()/isLower() or toUpper()/toLower() etc helper function provided by the language.  Your choice of language.

2. Consider this statement:
   ```
   $a = implode(',',array_map(function($b,$c) {
     return str_replace(array('-','_',','), '', $b) . "x{$c}";
   },array_keys($d),$d));
   ```
   a) what language is it written in?
   
   b) at the point when this statement is executed, which (if any) pre-existing variable(s) does this statement use or rely on?
   
   c) after this statement has executed, which (if any) variable(s) have been initialised or modified by the statement?
   
   d) taking your answer from b), give simple example value(s) for each used/relied-upon variable.  There is not a single correct answer, rather you should make an educated guess based on your interpretation of what the statement is doing.
   
   e) what would be the output or effect of the statement, if you used your example value(s) from d) ?
   
   f) describe what is happening in this statement

3. Write a function in Go which returns the top two most frequent numbers from a list, in order of frequency first. For example:
   ```
   Given the list [1, 3, 3, 5, 5, 6, 6, 5, 3, 3]
   It should return [3, 5]
   ```

4. Go to one of the Koala websites (au.koala.com, jp.koala.com)
   a) can you find our Shopify Storefront API key?  If so, what is it?
   b) based on what you found in a), is this an acceptable state-of-affairs for a modern eCommerce website?  Why or why not?

5. What will this PHP statement print?
   ```
   echo implode(' = ',['9 times 5','4' + '5']);
   ```

6. Using the students array below, write a javascript function to return an object containing:

  - The name of the class as the key.
  - The total attended lessons for each class.
  - The average amount of attended lessons for each class.

`students.json`

```json
{
  "students": [
    {
      "name": "Lulu Gearside",
      "class": "art",
      "attended": 35
    },
    {
      "name": "Matthew Milham",
      "class": "art",
      "attended": 11
    },
    {
      "name": "Dany Dufner",
      "class": "biology",
      "attended": 12
    },
    {
      "name": "Jeremy Doyle",
      "class": "biology",
      "attended": 3
    },
    {
      "name": "Tim O'Connor",
      "class": "biology",
      "attended": 10
    },
    {
      "name": "Charlie Wang",
      "class": "french",
      "attended": 12
    }
  ]
}
```

Expected output:

```js
{
  "art": {
    "total": 46,
    "average": 23,
  },
  "biology": {
    "total": 25,
    "average": 8,
  },
  "french": {
    "total": 12,
    "average": 12,
  },
}
```
