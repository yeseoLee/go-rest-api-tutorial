### DataSource
- DB 커넥션 획득하는 방법을 추상화 
- 애플리케이션 -> 데이터소스 -> 커넥션 풀 

```java
public interface DataSource {
    Connection getCOnnection() throws SQLException;
}
```

``` go 
type DataSource interface {
	getConnection() (Connection, error)
}
```