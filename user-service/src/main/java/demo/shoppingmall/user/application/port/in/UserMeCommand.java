package demo.shoppingmall.user.application.port.in;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class UserMeCommand {
    private final String userId;
}
