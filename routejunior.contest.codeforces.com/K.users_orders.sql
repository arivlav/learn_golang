SELECT users.id as id, users.name FROM orders
left JOIN users ON users.id=orders.user_id
group by users.id, users.name
having count(users.id)>0
order by users.name ASC, users.id ASC;