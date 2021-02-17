## Synchronizer

- Sincronizar é o ato de coordenar as atividades de duas ou mais threads. 

- Quando duas ou mais threads precisam acessar um recurso compartilhado, ou apenas uma thread pode acessar 
  o recurso por vez. 

- Por isso existe essa ideia de sincronizar as threads para que as operações de uma não interfira na outra.

- No Java é usado a palavra **synchronized** em métodos ou em blocos de códigos para realizar essa sincronização 
  de um método que esteja compartilhado por duas ou mais threads.

- Um exemplo da implementação disso está em `/threads/src/Calculadora`.
