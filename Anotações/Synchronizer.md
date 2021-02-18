## Synchronizer

- Sincronizar é o ato de coordenar as atividades de duas ou mais threads. 


- Quando duas ou mais threads precisam acessar um recurso compartilhado, ou apenas uma thread pode acessar 
  o recurso por vez. 


- Por isso existe essa ideia de sincronizar as threads para que as operações de uma não interfira na outra.


- No Java é usado a palavra **synchronized** em métodos ou em blocos de códigos para realizar essa sincronização 
  de um método que esteja compartilhado por duas ou mais threads.


- Um exemplo da implementação disso está em `/threads/src/Calculadora`.

## Threads: notify, wai e notify

- Uma thread A está executando dentro de um método sincronizado e precisa acessar um recurso R que 
  no momento está indisponível.

  
- Se a thread A ficar esperando por R, irá bloquear o objeto, impedindo que outras threads acessem o mesmo.


- Nesse caso a melhor solução para não causar problemas é liberar temporariamente o controlo do objeto 
  permitindo que outra thread seja executado. 


- Um exemplo é que duas pessoas precisam usar uma caneta. Para que uma não fique esperando até que a outra termine, 
  quem está usando pode disponibilizar a caneta e entrar em modo de espera. Quando quem estiver usando parar, 
  deve notificar a outra pessoa que a caneta está disponível. 

### Formas de aplicar isto em Java

- **wait**: bloqueia a execução da thread temporariamente, ou seja, coloca a thread em modo de espera.
A thread fica em modo de espera até que seja notificado. 


- **notify**: notifica uma thread que estava esperando, ou seja, retorna a execução da thread.


- **notifyAll**: notifica todas as threads, e a que tem prioridade mais alta ganha acesso ao objeto.

- Um exemplo da aplicação disso está em `/threads/src/TiqueTaque`.