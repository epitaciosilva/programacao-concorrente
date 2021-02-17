package Basico;

// Simple thread in Java
public class MyThread extends Thread {

    private String name;
    private int time;

    public MyThread(String name, int time) {
        this.name = name;
        this.time = time;
        this.start();
    }

    public void run() {
        System.out.println(this.name);
        try {
            for(int i = 0; i < 6; i++) {
                System.out.println(this.name +  ": " + i);
                Thread.sleep(time);
            }
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
