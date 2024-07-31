package demo.shoppingmall.user.adapter.out.persistence;

import demo.shoppingmall.user.domain.User;
import org.springframework.stereotype.Component;

@Component
public class UserMapper {
    User mapToDomain(UserJpaEntity userJpaEntity) {
        return User.generate(new User.UserId(userJpaEntity.getUserId()), new User.Name(userJpaEntity.getName()));
    }
}
