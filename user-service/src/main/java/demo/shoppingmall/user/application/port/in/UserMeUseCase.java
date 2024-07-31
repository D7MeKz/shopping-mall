package demo.shoppingmall.user.application.port.in;

import demo.shoppingmall.user.domain.User;

public interface UserMeUseCase {
    User userMe(UserMeCommand command);
}
