package demo.shoppingmall.user.adapter.in.web;

import demo.shoppingmall.user.application.port.in.UserMeCommand;
import demo.shoppingmall.user.application.port.in.UserMeUseCase;
import demo.shoppingmall.user.domain.User;
import demo.shoppingmall.user.utils.WebAdapter;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@WebAdapter
@RestController
@RequiredArgsConstructor
public class UserMeController {
    private final UserMeUseCase userMeUseCase;
    @GetMapping("/users/{userId}")
    public ResponseEntity<User> userMe(String userId) {
        UserMeCommand userMeCommand = UserMeCommand.builder().userId(userId).build();
        return ResponseEntity.ok(userMeUseCase.userMe(userMeCommand));
    }
}
