## AGENDA

> Basics of Golang
> Basics of MongoDB
> Basics of React.js
    


### Tasks

1. Create REST API in Golang using MongoDB as Database

2. Integrate API in React.js and render books in UI

 Mock this https://koenig-media.raywenderlich.com/uploads/2018/12/end_product-1.png UI



## 03/05/2020

 ## Getting Started with MongoDB

 #### Install MongoDB on Mac/Windows/Linux

    Download link: https://www.mongodb.com/download-center/community
    Document link: https://docs.mongodb.com/manual/administration/install-community/

#### Start Mongod service

    `brew services start mongodb-community@4.2`
    `brew services stop mongodb-community@4.2`
    `mongod --config /usr/local/etc/mongod.conf --fork`
    `ps aux | grep -v grep | grep mongod`

#### Start MongoDB

    `mongo`

#### CMD's


1. Show databases

    `show dbs`

2. Create Database

    `use stores`

3. Use Database

    `use store`
    
4. Drop Database

    `db.dropDatabase()`

5. Check wich DB

    `db`

6. Create Collections

    `db.createCollection('products')`
    
7. Show collections

    `show collections`

8. Insret query

    ```
    db.products.insert({
        title: 'Apple',
        details: 'More detils',
        grade: ['A', 'B', 'C', 'D'],
        tags: ['Apple', 'fruites', 'juice'],
        address:{
            pincode: 1234,
            area: "example_area"
        },
        farmers: [{
            name: 'person1',
            phno: 1234345,
        },
        {
            name: 'person2',
            phno: 12233333,
        }
        ],
        date: Date()
    })
    ```

    ```
    db.products.insert({
        title: 'Car',
        color: 'Red'
    })
    ```

9. Insert many query

    ```
    db.products.insertMany([
        {
        title: 'Mango',
        details: 'More detils',
        grade: ['A', 'B', 'C', 'D'],
        tags: ['Mango', 'fruites', 'juice'],
        address:{
            pincode: 1234,
            area: "example_area"
        },
        farmers: [{
            name: 'person1',
            phno: 1234345,
        },
        {
            name: 'person2',
            phno: 12233333,
        }
        ],
        date: Date()
    },
     {
        title: 'Apple',
        details: 'More detils',
        grade: ['A', 'B', 'C', 'D'],
        tags: ['Apple', 'fruites', 'juice'],
        address:{
            pincode: 1234,
            area: "example_area"
        },
        farmers: [{
            name: 'person1',
            phno: 1234345,
        },
        {
            name: 'person2',
            phno: 12233333,
        }
        ],
        date: Date()
    },
    ])

    ```
10. Query/Find/READ Data

    `db.products.find().pretty()`

11. Find particular

    `db.products.find({title: 'Apple'})`

12. Update

    ```
    db.products.update({title: 'Mango'}, 
        { $set: {grade:['A', 'B', 'C', 'D']}, 
        $push: { address:{ 
                pincode: 1234,
                area: "example_area"
            },
        farmers: [{
            name: 'person1',
            phno: 1234345,
        },
        {
            name: 'person2',
            phno: 12233333,
        }
        ]
        }})
    ```

13. UPDATE : Remove from array

    `db.products.update({title: 'Mango'}, {$pull: {'farmers': {'phno': 12233333}}})`

14. UPDATE : Remove feild from document

    `db.products.update({title: 'Apple'}, {$unset: {'details': 1}})`

    `db.products.update({title: 'Mango'}, {$unset: {'details': ""}})`



