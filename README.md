# volume-assignment
* This is typical graph problem, so used depth first search to solve this and built an api around it. 
* Time complexity of the problem is O(N)
* I spent roughly 2-3 hours to come up with base solution. Later spent some more time making it more performant and robust and realized im over complicating it and dropped the idea. 
* Below are the instruction on how to use the API.

**calculate**
----
  Takes a list of flights as input and returns the final path of the user

* **URL**

  [localhost:800/calculate](http://localhost:8080/calculate)

* **Method:**

  `POST`
* **Data Params**

 list of flights in text, for example [["SFO", "EWR"]]

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `["SFO", "EWR"]`
 
* **Error Response:**

  * **Code:** 500 Bad request <br />

* **Sample Call:**

  ```
curl -X POST -H "Content-Type: application/json" -d "[ [ \"IND\",\"EWR\"], [\"SFO\",\"ATL\"], [\"GSO\",\"IND\"], [\"ATL\",\"GSO\"] ]" http://localhost:8080/calculate



  ```
