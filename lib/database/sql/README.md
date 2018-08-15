# DB

1. DB is a database handle representing a pool of zero or more underlying connections. DB是一个句柄，用于管理底层的连接
2. The sql package creates and frees connections automatically. Mysql连接对上层调用是透明的
3. It's safe for concurrent use by multiple goroutines. 协程安全
4. Once DB.Begin is called, the returned Tx is bound to a single connection.
   Once Commit or Rollback is called on the transaction, that transaction's 
   connection is returned to DB's idle connection pool.
5. package database/sql 包里面有数据类型DB, Stmt, Tx 三种类型，他们都支持Exec、Stmt、
Query、QueryRow操作

- Open Conn and Idle Conn

The db pool may contain 0 or more idle connections to the database.
These were connections that were made, used, and rather than closed, were kept around for future use.
The number of these we can keep around is MaxIdleConns

When you request **one of these idle connections**, it becomes an Open connection, available for you to use.
The number of these you can use is MaxOpenConns.

```
MaxIdleConns <= MaxOpenConns
```

>
> Keeping an idle connection alive is not free
> It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.


## Tx

1. A transaction must end with a call to Commit or Rollback.
2. The statements prepared for a transaction by calling the transaction's    Prepare or Stmt methods are closed by the call to Commit or Rollback.

## [Locking Reads](https://dev.mysql.com/doc/refman/5.5/en/innodb-locking-reads.html)

If you query data and then insert or update related data within the same transaction, the regular SELECT statement does not give enough protection. Other transactions can update or delete the same rows you just queried.

提供了两种方法：

1. ```SELECT ... LOCK IN SHARE MODE```
2. ```SELECT ... FOR UPDATE```

>
> 语句也是要放在一个事务内部的

## [Tx 为什么要提供Query相关的方法？]https://stackoverflow.com/questions/1976686/is-there-a-difference-between-a-select-statement-inside-a-transaction-and-one-th

```
the one inside the transaction can see changes made by other previous Insert/
Update/delete statements in that transaction, A Select statement outside the 
transaction cannot....

If all you are asking about is what the Isolation Level does, then understand - 
that all Select Statements (hey, all statements of any kind), - are in a 
transaction. The only difference between one that is explicitly in a transaction 
and one that is standing on it's own is that the one that is standing alone 
starts it's transaction immediately before it executes it, and commits or roll 
back immediately after it executes,

whereas the one that is explicitly in a transaction can, (because it has a Begin 
Transaction statement) can have other statements (inserts/updates/deletes, 
whatever) occcurring within that same transaction, either before or after that 
Select statement.

So whatever the isolation level is set to, both selects (inside or outside an 
explicit transaction) will nevertheless be in a transaction which is operating 
at that isolation level.
```
