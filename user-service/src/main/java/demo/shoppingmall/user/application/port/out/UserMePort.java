package demo.shoppingmall.user.application.port.out;

import demo.shoppingmall.user.domain.User;

public interface UserMePort {
    User getUser(User.UserId userId);
}
