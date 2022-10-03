# Gerador de JSON de Demonstrativos Financeiros do BACEN

Este projeto tem como objetivo gerar um arquivo no formato JSON (_JavaScript Object Notation_) de acordo com as normas definidas pelo BACEN (Banco Central do Brasil) para Demonstrações Financeiras em JSON – Documento 90x1 ([Referência](https://www.bcb.gov.br/content/estabilidadefinanceira/Documents/Leiaute_de_documentos/CDSFN/Central_de_Demonstracoes_Financeiras-Leiautes_informacoes_tecnicasJSON.pdf)).

# Da entrada de dados

Esse programa deverá ser executado na mesma pasta onde se encontram os seguintes arquivos CVSs:

- `balanco.csv` - Referente ao **Balanço Patrimonial**;
- `dre.csv` - Referente à **Demonstração do Resultado (Demonstração de Sobras ou Perdas)**;
- `dra.csv` - Referente à **Demonstração do Resultado Abrangente**;
- `caixa.csv` - Referente à **Demonstração dos Fluxos de Caixa**;
- `dmpl.csv` - Referente à **Demonstração das Mutações do Patrimônio Líquido**

Estes arquivos devem ser gerados a partir da "[Planilha Template Demonstrativo Financeiro para JSON](https://docs.google.com/spreadsheets/d/1CtAKQ6g-tetm5wqZNW2sewtqI2fBZxIfjp-3cVRuX84/edit?usp=sharing)", visto que ela possui um padrão esperado pelo programa. Nela estão contidas maiores informações sobre o seu uso.

# Da execução do programa

TBD

# Da saída de dados

Após a execução do programa, este irá gerar um arquivo denominado `resultado.json`, que conterá as Demonstrações Financeiras, no formato JSON, no padrão Documento 90x1.
