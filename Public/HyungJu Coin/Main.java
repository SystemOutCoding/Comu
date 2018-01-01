/**
 * HyungJu Coin(...)의 메인 클래스가 여기있어요
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

import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.ObjectOutputStream;
import java.rmi.server.ExportException;
import java.util.Base64;


public class Main {
    public static int  byteToint(byte[] arr){
        return (arr[0] & 0xff)<<24 | (arr[1] & 0xff)<<16 | (arr[2] & 0xff)<<8 | (arr[3] & 0xff);
    }

    public static BlockChain createBlock(String previousblockhash, String merklehash, long bits, long nonce, Transaction transaction[]){

        String version = "HyungJu Coin";

        BlockChain blockChain = new BlockChain();
        BlockHeader header = new BlockHeader(version, previousblockhash, merklehash, bits, nonce);
        //header 직렬화 시작
        String headerHash=null;
        ObjectMapper mapper = new ObjectMapper();
        try{


           headerHash = SHA256.encrypt(SHA256.encrypt(mapper.writeValueAsString(header)));
        }catch (Exception e){

        }
        blockChain.hash = headerHash;
        blockChain.transaction = transaction;
        blockChain.header = header;


        return blockChain;
    }

    private static String muckleTree(String transaction[]){

          if(transaction.length == 1) {
              //muckleTree 루트가 계산이 완료됨.
              return transaction[0];
          }
          // hash된 거래의 1/2의 크기로 String 배열을 만듦
          String combinedTransaction[] = new String[transaction.length/2];
          int count=0; //combinedTransaction에 두 plaintext가 더해져 만들어진 transaction의 hash를 저장하는데 필요한
          //index 값
          for(int i=0;i<transaction.length;i= i+2){

              combinedTransaction[count]=SHA256.encrypt(SHA256.encrypt(transaction[i]+transaction[i+1]));
             // System.out.println(transaction[i]+"와"+transaction[i+1]+"는"+combinedTransaction[count]);
              count++;
          }

          return muckleTree(combinedTransaction);

    }
    private static String[] transactionSerialization(Transaction transaction[]){


        //거래 Class 직렬화 시작
        String Transactions[] = new String[transaction.length];
        for(int i=0;i< transaction.length; i++){
            ObjectMapper mapper = new ObjectMapper();
            try{

                Transactions[i] = mapper.writeValueAsString(transaction[i]);
               // System.out.println(Transactions[i]);
                // base 64 변환된 데이터를 hash 암호화 함.
                Transactions[i] = SHA256.encrypt(Transactions[i]);
            }catch (Exception e){

            }


        }

        return Transactions;

    }
    public static void main(String[] args){
        Transaction[] transaction = new Transaction[4];
        transaction[0] = new Transaction(0, "나","너", 100,"돈 보낸다");
        transaction[1] = new Transaction(1, "너","나", 100,"돈 갚는다");
        transaction[2] = new Transaction(1, "너","나", 100,"돈 갚는다");
        transaction[3] = new Transaction(1, "너","나", 100,"돈 갚는다");
        //System.out.println(muckleTree(transactionSerialization(transaction)));

        //Genesis Block 생성
        BlockChain blockChain = createBlock("Genesis",muckleTree(transactionSerialization(transaction)),1,134, transaction);

        int i=0;
        while(true){

            ObjectMapper mapper = new ObjectMapper();
            blockChain.header.nonce = i;
            try{

                String headerHash = SHA256.encrypt(mapper.writeValueAsString(blockChain.header));
               // if(headerHash)
                char[] hh = new char[headerHash.length()];
                char[] bc = new char[blockChain.hash.length()];
                hh = headerHash.toCharArray();
                bc = blockChain.hash.toCharArray();

                if(hh[0]<bc[0]){
                    System.out.println("ok");
                }

            }catch (Exception e){

            }
            i++;
         //   if(temp.hash< blockChain.hash){
           // }
        }

    }
}
