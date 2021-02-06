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

Answer:


```
doSomething
|-----------------|
                  doSomethingElse(undefined)
                  |------------------|
                                     finalHandler(undefined)
                                     |------------------|
```

#### Puzzle #3

```js
doSomething()
  .then(doSomethingElse())
  .then(finalHandler);
```


Answer:

```
doSomething
|-----------------|
                  doSomethingElse(undefined)
                  |------------------|
                                     finalHandler(resultOfDoSomething)
                                     |------------------|
```

#### Puzzle #4

```js
doSomething()
  .then(doSomethingElse)
  .then(finalHandler);
```


Answer:

```
doSomething
|-----------------|
                  doSomethingElse(resultOfDoSomething)
                  |------------------|
                                     finalHandler(resultOfDoSomethingElse)
                                     |------------------|
```

## Quick challenges

### Shell/CLI
1. What is happening in these Linux commands?  Describe as much as you know about what each symbol means and the effect it will have in the execution of the command, and the named programs.

   a) `(for i in {1..100}; do echo $i; done;) | grep 3 | grep -v 1 | paste -s -d+ - | bc`

   Breakdown explain of every command:

   Explain:
   * `(for i in {1..100}; do echo $i; done;) ` : do the loops from 1 to 100 and print it out using `echo`
   * `|`: execute the preceding statement and connect its stdout the to stdin of the statement which follows. 
   So `(for i in {1..100}; do echo $i; done;) | grep 3 ` mean execute the loop first and then execute the `grep 3` command
   * `grep 3`: The grep command is used to search text. In this case, it will search for all text that include `3`.
   Effect:
   ```
    3
    13
    23
    30
    31
    32
    33
    ...
   ```
   * `grep -v 1`: the `-v` in `grep` here to select non-matching lines so in this case it will select all the text that not include `1` from the previes command value (`grep 3`). 
   Effect:
   ```
    3
    23
    30
    32
    33
    34
    35
    ...
   ``` 
   * `paste -s`: merge all value into a single line.
   Effect:
   ```
   3	23	30	32	33	34 ...
   ```
   * `paste -s -d+`: merge all value into single line and join them with `+`
   Effect:
   ```
   3+23+30+32+33+34...
   ```
   * `paste -s -d+ -`: The hyphens (`-`) at the end just combining 1 consecutive lines 

   * `bc`: used for command line calculator, basically it will calculate the paste value and print it to the screen which is 748
   ```
   3+23+30+32+33+34... = 748
   ```
   
   b) `[ ! -f /var/lock/myscript.lock ] && touch /var/lock/myscript.lock && (yum -y update >> /var/log/mylog.log 2>&1; ) && rm -f /var/lock/myscript.lock`


2. There is a directory, containing a large tree of subdirectories and files.  Scattered throughout these files are Australian phone numbers, and we want to harvest them – we want to end up with a simple list of the phone numbers.

   a) Write a regular expression to match Australian phone numbers.  The numbers will be in a mixture of the forms 02xxxxxxxx and +612xxxxxxxx, and there will also be the common usage of hyphens, spaces, and parentheses, so all of those common possibilities must be supported. eg. 02 xxxx xxxx, (02) xxxx-xxxx, +61 2 xxxx xxxx, +61 02 xxxxxxxx, +61 (0)2 xxxx-xxxx

   Assumption:
   * 61 or +61 is a replacement for 0
   * we only look for the `02` (ex: 02 2379 5821) format phone number (04 2379 5821 will not be accepted)

   ```
    ^[(+]{0,2}(61|0){1}[) ]*(\(?0?\)?)*(2\)?)+[0-9\.\- ]+
   ```

    Explain:

    * `^`: start with (remove this if you want to find it in all text, ex: "Phone number: `0423795821`" it will match the phone number only but not the text)
    * `[(+]`: any character within this set which is `(` `+` 
    * `{0,2}`: that match 0 to 2 
    * `(61|0){0,1}:` match 0 to 1 either 61 or 0 (used to migrate phone number before using regex and I saw people put phone number without the `0` such as 423795821)
    Ex: `+61`, `0`,`(0`, `(61`, `61` or `(+61`, etc
    * `[) ]*`: match 0 or more the close parentheses or a space
    * `(\(?0?\)?)*`: match optional `(`, `0`, `)`
    * `(2\)?)+`: match at least 1 the `2` with optional `?`. 
    * `[0-9\.\- ]+`: match the set of number, `.` (something I see phone in this phone format 0423.795.821), and `-`

   b) Write an example phone number, for each specific phone number format that your regex would match.
   ```
    +612xxxxxxxx        : +61223795821
    +612.xxxx.xxxx      : +612.2379.5821
    02XXXXXXXX          : 0223795821
    02 xxxx xxxx        : 02 2379 5821
    (02) xxxx-xxxx      : (02) 2379 5821
    +61 2 xxxx xxxx     : +61 2 2379 5821
    +61 02 xxxxxxxx     : +61 02 2379 5821
    +61 (0)2 xxxx-xxxx  : +61 (0)2 2379 5821
   ```

   c) Imagine that the full path of the directory is
   > /var/www/site1/uploads/phnumbers/
   
   Write a single-line or simple command that you could run from the shell (ideally Linux), to apply this regular expression to the files in the directory tree, and result in the simple list of phone numbers, one phone number per line.


### Software development

1.	Write a function/method/subroutine to determine if a string starts with an upper-case letter A-Z, *without* using any sort of isUpper()/isLower() or toUpper()/toLower() etc helper function provided by the language.  Your choice of language.

Language: Javascript

There are two methods here because I don't whether regex can be used here so I provided two answers: one using regex and one just purely javascript

Assumption:
* The string is follow this format: "Hello world" -> In this case the `H` is the first letter and it upper case so it true
* "hello World" -> In this case the `h` is the first letter and it not upper case so it false

```
const is_uppercase_regex = (string) => /^[A-Z]*$/.test(string[0]);

const is_uppercase = (string) => {
    let value = false;

    const alphabet = 'QWERTYUIOPASDFGHJKLZXCVBNM'

    for (let i = 0; i < alphabet.length; i++) {
        if(string[0] === alphabet[i]) {
            value = true
            break;
        }
      }

    return value;
}
```

2. Consider this statement:
   ```
   $a = implode(',',array_map(function($b,$c) {
     return str_replace(array('-','_',','), '', $b) . "x{$c}";
   },array_keys($d),$d));
   ```
   a) what language is it written in?

   ```
   PHP
   ```
   
   b) at the point when this statement is executed, which (if any) pre-existing variable(s) does this statement use or rely on?

   `It will execute when the an array $d is defined (ex: $d = array(1,2,3)). In another words, the statement is rely on variable $d and it type Array`
   
   c) after this statement has executed, which (if any) variable(s) have been initialised or modified by the statement?
   
   d) taking your answer from b), give simple example value(s) for each used/relied-upon variable.  There is not a single correct answer, rather you should make an educated guess based on your interpretation of what the statement is doing.

   * `$d`: `array( "hei-ght" => "100", "wi-dth" => "200" )`
   * `array_keys($d)`: `array( 0 => "hei-ght" , 1 => "wi-dth")`
   * `$b`: has value of second args which is `array_keys($d)`, first iteration will be `hei-ght` and second iteration will be `wi-dth`
   * `$c`: has value of third args which is `$d`, fourth iteration will be `100` and third iteration will be `200`
   * `str_replace(array('-','_',','), '', $b) . "x{$c}";`: will return `heightx100` and `widthx200`
   * `array_map`: result will be an array with value as `array( 0 => "heightx100", 1 => "widthx200")`
   * `$a`: will be the value of the join from `implode` which is a string: `heightx100,widthx200`
   
   e) what would be the output or effect of the statement, if you used your example value(s) from d) ?

   output: A string and it value is `heightx100,widthx200`
   
   f) describe what is happening in this statement

   * The function executes the `implode` to join the second args, which the result of `array_map` return

   * `array_map` take the first args as a callback function to run for each element in each array and number of arrays accepted in this case there are two arrays passed to the the callback. The array of `array_keys($d)` (array_keys will return an array of all the keys of array `$d`) and the actual array `$d`

   Example: if there the `$d` value is: `$d = array( "height" => "100", "width" => "200" )`
   then the array_map will look like this:
    ```
    array_keys ( callback, ["height", "width"] , ["height" => "100", "width" => "200"] )

    ```
   * Callback function accepts match the number of arrays passed to `array_map` in this case is two. 
   ```
    array_map(function($b,$c) {
      ...
    }
   ```
   from example above, $b has value of array `array_keys($d)` ('height' and 'width') and $c has value of array `$d` ('100' and '200')

   * `array_map` will return the value of `return str_replace(array('-','_',','), '', $b) . "x{$c}";`, where `str_replace` take the first args to search (it will search for `-`, `_`, `,`) and second args to replace it with `''` (empty value)

   and it will concatenate the value from `str_replace` to `"x${$c}"` where `${c}` with complex (curly) syntax will return `$c` value

   Example: so from the example above, the `$a` has value:

   ```
    [
      0 => heightx100, // First iteration (`height is the value of str_replace` + "x" + `100 is value of $c`)
      1 => widthx200   // Second iteration (`width is the value of str_replace` + "x" + `200 is value of $c`)
    ]
   ```

3. Write a function in Go which returns the top two most frequent numbers from a list, in order of frequency first. For example:
   ```
   Given the list [1, 3, 3, 5, 5, 6, 6, 5, 3, 3]
   It should return [3, 5]
   ```
  Assumption:
  * There can only be two most frequent numbers in the array (array like this `[2, 2, 1, 1, 3, 3]` with three frequency will not be accepted)

   Anwser:
   ```
    func twoMostFrequent(nums []int) []int {

        // Hashmap to hold the count of the element
        hash := make(map[int]int)

        count_most_freq := 0

        count_second_freg := 0

        most := 0

        second := 0

        results := []int{}

        // Just return empty if provided array is empty
        if len(nums) == 0 {
          return results
        }

        for _, value := range nums {

          // For checking if the element exist in the hash
          _, ok := hash[value]

          if !ok {
            hash[value] = 0
          }
          hash[value] = hash[value] + 1
        }

        for key, value := range hash {
          if value > count_most_freq {
            second = most
            count_second_freg = count_most_freq

            most = key
            count_most_freq = value
          } else if count_second_freg < count_most_freq && value > count_second_freg {
            second = key
            count_second_freg = value
          }
        }

        results = []int{most, second}

        return results
      }
   ```

4. Go to one of the Koala websites (au.koala.com, jp.koala.com)
   a) can you find our Shopify Storefront API key?  If so, what is it?
   b) based on what you found in a), is this an acceptable state-of-affairs for a modern eCommerce website?  Why or why not?

5. What will this PHP statement print?
   ```
   echo implode(' = ',['9 times 5','4' + '5']);
   ```

   Anwser:

    implode: join array

    ['9 times 5','4' + '5'] there is two value in this array:

    * position 0: a string with value '9 times 5'
    * position 1: a int with value 9 (because in PHP it will parse 4 and 5 to int value when use `+`)

   it return a string below:

   ```
    '9 times 5 = 9'
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

Anwser

```
const data = require('./students.json');

const totalStudent = (students) => {

    let returnValue = [];

    // An object to capture the length of the class to calculate the average student
    let repeatter = {};

    students.map((student) => {

        let shouldPush = true;
        returnValue.map((value, index) => {
            if(value[student['class']]) {
                shouldPush = false;

                returnValue[index][student['class']]['total'] += student.attended;

                repeatter[student['class']] += 1;

                returnValue[index][student['class']]['average'] =  
                Math.round(
                    returnValue[index][student['class']]['total'] / repeatter[student['class']]
                )
            } 
        })

        if(shouldPush) {

            repeatter[student['class']] = 1;

            returnValue.push({
                [student['class']]: {
                    total: student.attended,
                    average: Math.round(student.attended/repeatter[student['class']])
                }
            })
        }
    });

    return returnValue;
}


// Assumption
// attended: Integer
// students: Array
console.log(totalStudent(data.students));
```