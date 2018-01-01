/**
 * SHA256 암호화 코드에요
 * http://blog.kindler.io/java-encrypt/ 에서 참고해왔으나
 * 이분의 저작권(?) 정책을 찾을수가 없어서 메인의 카피라이트와 같게 설정하겠습니다.
 *
 * 카피라이트요?
 * GNU General Public License v3.0 쓸래요
 *
 * 아참, 각각 변수 메소드들의 대한 설명은
 * 이거 개발할때 안써서 못적었어요
 * 양해바래요
 *
 * 그리고 이 코인(..)은 아직 제작중이에요
 * P2P 통신도 구현되있지않고
 * 개발자의 실력부족으로 hash의 크기비교도 못하고 있는 실정이에요
 * 정말 미안해요
 *
 *
 *
 */

import java.security.MessageDigest;

public class SHA256 {
    public static String encrypt(String planText) {
        try{
            MessageDigest md = MessageDigest.getInstance("SHA-256");
            md.update(planText.getBytes());
            byte byteData[] = md.digest();

            StringBuffer sb = new StringBuffer();
            for (int i = 0; i < byteData.length; i++) {
                sb.append(Integer.toString((byteData[i] & 0xff) + 0x100, 16).substring(1));
            }

            StringBuffer hexString = new StringBuffer();
            for (int i=0;i<byteData.length;i++) {
                String hex=Integer.toHexString(0xff & byteData[i]);
                if(hex.length()==1){
                    hexString.append('0');
                }
                hexString.append(hex);
            }

            return hexString.toString();
        }catch(Exception e){
            e.printStackTrace();
            throw new RuntimeException();
        }
    }


}
