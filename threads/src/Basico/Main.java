package Basico;

public class Main {

    public static void main(String[] args) {
//        MyThread thread1 = new MyThread("Thread 1", 600);
//        MyThread thread2 = new MyThread("Thread 2", 900);

        ThreadWithRunnable thread1 = new ThreadWithRunnable("Thread 1", 500);
        ThreadWithRunnable thread2 = new ThreadWithRunnable("Thread 2", 500);
        ThreadWithRunnable thread3 = new ThreadWithRunnable("Thread 3", 500);
        Thread t1 = new Thread(thread1);
        Thread t2 = new Thread(thread2);
        Thread t3 = new Thread(thread3);

        t2.start();
        t1.start();
        t3.start();

        try {
            t1.join();
            t2.join();
            t3.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("Program has finished!");
    }

}
