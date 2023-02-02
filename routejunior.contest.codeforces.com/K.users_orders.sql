SELECT orders.user_id, users.name FROM orders
JOIN users ON users.id=orders.user_id
group by orders.user_id, users.name
having count(orders.user_id)>0
order by users.name ASC, orders.user_id ASC;