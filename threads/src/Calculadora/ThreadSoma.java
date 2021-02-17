package Calculadora;

public class ThreadSoma implements Runnable {

    private String name;
    private int[]  nums;

    // recurso compartilhado
    private static Calculadora calc = new Calculadora();

    public ThreadSoma(String name, int[] nums) {
        this.name = name;
        this.nums = nums;
        new Thread(this, name).start();
    }

    @Override
    public void run() {
        System.out.println(this.name + " iniciado");

        int soma = calc.somaArray(this.nums);

        System.out.println("Resultado da soma: " + soma);
        System.out.println(this.name + " terminado");
    }

}
