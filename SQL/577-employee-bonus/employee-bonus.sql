/* Write your T-SQL query statement below */
SELECT e.name, b.Bonus from Employee as e
left join Bonus as b on e.empId = b.empId
where b.empid is null 
or b.bonus < 1000