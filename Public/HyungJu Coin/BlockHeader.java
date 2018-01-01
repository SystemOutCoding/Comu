/**
 * BlockHeader에요.
 * 여기에 블록의 설명이 담겨있고 나중에는 이안에 nonce로 채굴을 하게되겠죠
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
 * @author HyungJu
 *
 */

import java.io.Serializable;

public class BlockHeader implements Serializable {

    public String  version, previousblockhash, merklehash;
    public long time, bits, nonce;

    BlockHeader(String version, String previousblockhash, String merklehash, long bits, long nonce){
        time = System.currentTimeMillis() / 1000;
    }
}
