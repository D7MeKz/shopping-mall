package demo.shoppingmall.user.domain;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Value;

@Getter
@AllArgsConstructor(access = AccessLevel.PRIVATE)
public class User {

    private final String userId;
    private final String name;

    public static User generate(UserId userId, Name name) {
        return new User(userId.userId, name.name);
    }

    @Value
    public static class UserId {
        String userId;
        public UserId(String userId) {
            this.userId = userId;
        }

        public UserId(Long userId) {
            this.userId = String.valueOf(userId);
        }
    }

    @Value
    public static class Name {
        public Name(String name) {
            this.name = name;
        }
        String name;
    }
}
