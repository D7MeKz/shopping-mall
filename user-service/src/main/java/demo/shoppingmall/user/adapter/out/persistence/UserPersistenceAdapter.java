package demo.shoppingmall.user.adapter.out.persistence;

import demo.shoppingmall.user.application.port.out.UserMePort;
import demo.shoppingmall.user.domain.User;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;

@RequiredArgsConstructor
@Component
public class UserPersistenceAdapter implements UserMePort {
    private final UserRepository userRepository;
    private final UserMapper userMapper;
    @Override
    public User getUser(User.UserId userId) {
        return userMapper.mapToDomain(userRepository.getById(Long.parseLong(userId.getUserId())));
    }
}
