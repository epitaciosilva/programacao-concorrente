package Basico;

public class ThreadWithRunnable implements Runnable {

    private String name;
    private int time;

    public ThreadWithRunnable(String name, int time) {
        this.name = name;
        this.time = time;
    }

    @Override
    public void run() {
        try {
            for(int i = 0; i < 6; i++) {
                System.out.println(this.name + ": " + i);
                Thread.sleep(this.time);
            }
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println(name + " has finished!");
    }
}
