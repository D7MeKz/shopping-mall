package demo.shoppingmall.user.service;

import demo.shoppingmall.user.application.port.in.UserMeCommand;
import demo.shoppingmall.user.application.port.in.UserMeUseCase;
import demo.shoppingmall.user.application.port.out.UserMePort;
import demo.shoppingmall.user.domain.User;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

@RequiredArgsConstructor
@Component
@Transactional
public class UserMeService implements UserMeUseCase {
    private final UserMePort userMePort;
    @Override
    public User userMe(UserMeCommand command) {
//        return userMePort.getUser(new User.UserId(command.getUserId()));
        return User.generate(new User.UserId("1"), new User.Name("test"));
    }
}
