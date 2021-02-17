package Calculadora;

public class Calculadora {
    private int soma;

    public synchronized int somaArray(int[] array ) {
        soma = 0;

        for (int i : array) {
            soma += i;

            System.out.println("Executando a soma " + Thread.currentThread().getName() +
                    " somando o valor " + i + " com total de " + soma);

            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

        }
        return soma;
    }
}
