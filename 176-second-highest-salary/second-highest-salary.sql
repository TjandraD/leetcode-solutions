# Write your MySQL query statement below
SELECT MAX(DISTINCT salary) AS 'SecondHighestSalary'
FROM Employee
WHERE salary < (
    SELECT salary
    FROM Employee
    ORDER BY salary DESC
    LIMIT 1
);