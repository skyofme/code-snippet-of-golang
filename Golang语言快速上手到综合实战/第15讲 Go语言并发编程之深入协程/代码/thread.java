
public class thread {
	public static void main(String[] args){
		new Thread(new Test_thread()).start();
		new Thread(new Test_thread2()).start();
	}
}

class Test_thread implements Runnable{
	public void run() {
		  for (int i = 0; i < 100; i++) {
	            System.out.println("thread 1:" + i);
	      }	
	}
}

class Test_thread2 implements Runnable{
	public void run() {
		  for (int i = 100; i < 200; i++) {
	            System.out.println("thread 2:" + i);
	      }	
	}
}