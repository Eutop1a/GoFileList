package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func JudgeAccessToken(msg string) (string, bool) {

	//client, _ := rpc.DialHTTP("tcp", "localhost:65532")

	token := Token{
		TokenString: msg,
	}
	var result Result

	err := token.RPCParseToken(token, &result)

	//err := client.Call("Token.RPCParseToken", &token, &result)

	if err != nil || result.Status != "" {
		return "", false
	}

	return result.UserName, true
}

// RPCParseToken 解析token
func (token *Token) RPCParseToken(tokenString Token, result *Result) error {
	// 解析 JWT
	parsedToken, err := jwt.Parse(tokenString.TokenString, func(token *jwt.Token) (interface{}, error) {
		// 指定用于验证签名的密钥
		// 这个函数会在解析 JWT 时被调用
		return Secret, nil
	})
	// 处理解析过程中的错误
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				result.Status = "Token has expired"
				result.UserName = ""
				fmt.Println("Token has expired")
				return ve
			} else {
				result.Status = "Error parsing token:"
				result.UserName = ""
				fmt.Println("Error parsing token:", err.Error())
				return err
			}
		}
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// 获取声明中的具体数据
		username, err := claims["username"].(string)
		if err == true {
			//fmt.Println("username arrest success")
			result.UserName = username
		} else {
			// 处理类型断言失败的情况
			fmt.Println("username arrest error")
		}
		result.Status = ""
		// 使用解析后的数据进行相应的操作
		//fmt.Println("Username:", result.UserName)
		//fmt.Println("Role:", result.Role)
	} else {
		fmt.Println("Invalid token")
		result.Status = "Invalid token"
		result.UserName = ""
		//result.Role = ""
		return fmt.Errorf("invalid token")
	}
	return nil
}
