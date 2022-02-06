package Calculadora;

public class Main {

    public static void main(String[] args) {
        int[] array = {1,2,3,4,5};
        new ThreadSoma("Thread 1", array);
        new ThreadSoma("Thread 2", array);

    }

}

