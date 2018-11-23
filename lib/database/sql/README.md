## 数据库相关资料

- http://go-database-sql.org/
- database-sql
- github.com/go-sql-driver/mysql/

## Working with NULLs

1. Nullable columns are annoying and lead to a lot of ugly code. If you can, avoid them.
2. reasons to avoid nullable columns
    - There’s no sql.NullUint64 or sql.NullYourFavoriteType
    - If you need to define your own types to handle NULLs, you can copy the design of sql.NullString to achieve that.
    - `SELECT name, COALESCE(other_field, '') as otherField WHERE id = ?`
    - If `other_field` was NULL, `otherField` is now an empty string. This works with other data types as well

## Open Conn and Idle Conn

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

## Tx

1. Once DB.Begin is called, the returned Tx is bound to a single connection.
   Once Commit or Rollback is called on the transaction, that transaction's
   connection is returned to DB's idle connection pool.
2. A transaction must end with a call to Commit or Rollback.
3. The statements prepared for a transaction by calling the transaction's    Prepare or Stmt methods are closed by the call to Commit or Rollback.

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
