# Write your MySQL query statement below
SELECT subordinate.name AS 'Employee'
FROM Employee subordinate
JOIN Employee manager ON subordinate.managerId = manager.id
WHERE manager.salary IS NOT NULL AND manager.salary < subordinate.salary;