# Gerador de JSON de Demonstrativos Financeiros do BACEN

Este projeto tem como objetivo gerar um arquivo no formato JSON (_JavaScript Object Notation_) de acordo com as normas definidas pelo BACEN (Banco Central do Brasil) para Demonstrações Financeiras em JSON – Documento 90x1 ([Referência](https://www.bcb.gov.br/content/estabilidadefinanceira/Documents/Leiaute_de_documentos/CDSFN/Central_de_Demonstracoes_Financeiras-Leiautes_informacoes_tecnicasJSON.pdf)).

![image](https://github.com/oliveirafilipe/demonstracao-finaceira-bacen-json/assets/22033274/79f4a50c-ee09-4b7f-8873-cd0b8f9002ae)

# Da entrada de dados

Esse programa recebe arquivos CSVs como entrada para gerar um único arquivo de saída de formato JSON.

Estes arquivos de entrada devem ser gerados a partir da "[Planilha Template Demonstrativo Financeiro para JSON](https://docs.google.com/spreadsheets/d/1CtAKQ6g-tetm5wqZNW2sewtqI2fBZxIfjp-3cVRuX84/edit?usp=sharing)", visto que ela possui um padrão esperado pelo programa. Nela estão contidas maiores informações sobre o seu uso.

> ⚠️ **Atente-se** ao fato de que um CSV não carregará consigo nenhuma formatação de estilo (negrito, itálico, sublinhado) e que portanto as regras presentes na Planilha Template devem ser corretamente seguidas, para que o programa identifique Subitens.

Após corretamente preenchido, baixe cada planilha no formato CSV, agrupe-os em uma pasta e garanta que estes seguem a seguinte nomenclatura:

- `balanco.csv` - Referente ao **Balanço Patrimonial**;
- `dre.csv` - Referente à **Demonstração do Resultado (Demonstração de Sobras ou Perdas)**;
- `dra.csv` - Referente à **Demonstração do Resultado Abrangente**;
- `caixa.csv` - Referente à **Demonstração dos Fluxos de Caixa**;
- `dmpl.csv` - Referente à **Demonstração das Mutações do Patrimônio Líquido**

# Da execução do programa

1. Baixe [a última versão](https://github.com/oliveirafilipe/demonstracao-finaceira-bacen-json/releases) deste programa, leve em conta seu sistema operacional.
2. Copie o arquivo do programa para a mesma pasta onde estão os CSVs
3. Execute o programa:
   1. Se você está no **Windows**: Clique duas vezes para executar o programa
   2. Se você está no **Linux**: Execute via terminal `$ ./demonstracao-finaceira-bacen-json`
4. Leia atentamente as requisições do programa para que você informe corretamente as entradas de dados.
5. O programa lhe informará mensagens de erro e atenção, leia com atenção.
6. Se o programa executar com sucesso, será possível encontrar na pasta de execução um novo arquivo chamado `resultado.json`

# Da saída de dados

Após a execução do programa, este irá gerar um arquivo denominado `resultado.json`, que conterá as Demonstrações Financeiras, no formato JSON, no padrão Documento 90x1.
