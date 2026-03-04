/* Write your T-SQL query statement below */
SELECT Name from Employee 
where id in 
(
    Select managerId from Employee
    group by managerId 
    having count(*) > 4
)