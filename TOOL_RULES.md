# Here is the full list of rules that can be used for data extraction

*
    ```perl
    _[X:]
    ```
     Pass first X characters of the rest.
	 ### Example
	 Let the rest is
	 ```
	 1234abcd
	 ```

      Then the rule
	 ```perl
	 _[4:]
	 ```
	  will change the rest into
	  ```
	  abcd
	  ```
