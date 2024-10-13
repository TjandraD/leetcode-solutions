# Write your MySQL query statement below
SELECT DISTINCT Person.email
FROM Person
JOIN (
    SELECT email, COUNT(id) AS 'count'
    FROM Person
    GROUP BY email
) email_count ON email_count.email = Person.email
WHERE email_count.count > 1;