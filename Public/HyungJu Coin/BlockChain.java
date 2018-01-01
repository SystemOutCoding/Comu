/**
 * 여기는 BlockChain 클래스에요
 * 사실 한 Block이지만 편의상 블록체인이라고 했어요
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

public class BlockChain {
    String hash;
    BlockHeader header;
    Transaction transaction[];
}
