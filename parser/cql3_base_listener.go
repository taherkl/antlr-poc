// Generated from CQL3.g4 by ANTLR 4.7.

package parser // CQL3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCQL3Listener is a complete listener for a parse tree produced by CQL3Parser.
type BaseCQL3Listener struct{}

var _ CQL3Listener = &BaseCQL3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCQL3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCQL3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCQL3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCQL3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseCQL3Listener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseCQL3Listener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCQL3Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCQL3Listener) ExitStatement(ctx *StatementContext) {}

// EnterDml_statements is called when production dml_statements is entered.
func (s *BaseCQL3Listener) EnterDml_statements(ctx *Dml_statementsContext) {}

// ExitDml_statements is called when production dml_statements is exited.
func (s *BaseCQL3Listener) ExitDml_statements(ctx *Dml_statementsContext) {}

// EnterDml_statement is called when production dml_statement is entered.
func (s *BaseCQL3Listener) EnterDml_statement(ctx *Dml_statementContext) {}

// ExitDml_statement is called when production dml_statement is exited.
func (s *BaseCQL3Listener) ExitDml_statement(ctx *Dml_statementContext) {}

// EnterCreate_keyspace_stmt is called when production create_keyspace_stmt is entered.
func (s *BaseCQL3Listener) EnterCreate_keyspace_stmt(ctx *Create_keyspace_stmtContext) {}

// ExitCreate_keyspace_stmt is called when production create_keyspace_stmt is exited.
func (s *BaseCQL3Listener) ExitCreate_keyspace_stmt(ctx *Create_keyspace_stmtContext) {}

// EnterAlter_keyspace_stmt is called when production alter_keyspace_stmt is entered.
func (s *BaseCQL3Listener) EnterAlter_keyspace_stmt(ctx *Alter_keyspace_stmtContext) {}

// ExitAlter_keyspace_stmt is called when production alter_keyspace_stmt is exited.
func (s *BaseCQL3Listener) ExitAlter_keyspace_stmt(ctx *Alter_keyspace_stmtContext) {}

// EnterDrop_keyspace_stmt is called when production drop_keyspace_stmt is entered.
func (s *BaseCQL3Listener) EnterDrop_keyspace_stmt(ctx *Drop_keyspace_stmtContext) {}

// ExitDrop_keyspace_stmt is called when production drop_keyspace_stmt is exited.
func (s *BaseCQL3Listener) ExitDrop_keyspace_stmt(ctx *Drop_keyspace_stmtContext) {}

// EnterUse_stmt is called when production use_stmt is entered.
func (s *BaseCQL3Listener) EnterUse_stmt(ctx *Use_stmtContext) {}

// ExitUse_stmt is called when production use_stmt is exited.
func (s *BaseCQL3Listener) ExitUse_stmt(ctx *Use_stmtContext) {}

// EnterCreate_table_stmt is called when production create_table_stmt is entered.
func (s *BaseCQL3Listener) EnterCreate_table_stmt(ctx *Create_table_stmtContext) {}

// ExitCreate_table_stmt is called when production create_table_stmt is exited.
func (s *BaseCQL3Listener) ExitCreate_table_stmt(ctx *Create_table_stmtContext) {}

// EnterAlter_table_stmt is called when production alter_table_stmt is entered.
func (s *BaseCQL3Listener) EnterAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// ExitAlter_table_stmt is called when production alter_table_stmt is exited.
func (s *BaseCQL3Listener) ExitAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// EnterAlter_table_instruction is called when production alter_table_instruction is entered.
func (s *BaseCQL3Listener) EnterAlter_table_instruction(ctx *Alter_table_instructionContext) {}

// ExitAlter_table_instruction is called when production alter_table_instruction is exited.
func (s *BaseCQL3Listener) ExitAlter_table_instruction(ctx *Alter_table_instructionContext) {}

// EnterDrop_table_stmt is called when production drop_table_stmt is entered.
func (s *BaseCQL3Listener) EnterDrop_table_stmt(ctx *Drop_table_stmtContext) {}

// ExitDrop_table_stmt is called when production drop_table_stmt is exited.
func (s *BaseCQL3Listener) ExitDrop_table_stmt(ctx *Drop_table_stmtContext) {}

// EnterTruncate_table_stmt is called when production truncate_table_stmt is entered.
func (s *BaseCQL3Listener) EnterTruncate_table_stmt(ctx *Truncate_table_stmtContext) {}

// ExitTruncate_table_stmt is called when production truncate_table_stmt is exited.
func (s *BaseCQL3Listener) ExitTruncate_table_stmt(ctx *Truncate_table_stmtContext) {}

// EnterCreate_index_stmt is called when production create_index_stmt is entered.
func (s *BaseCQL3Listener) EnterCreate_index_stmt(ctx *Create_index_stmtContext) {}

// ExitCreate_index_stmt is called when production create_index_stmt is exited.
func (s *BaseCQL3Listener) ExitCreate_index_stmt(ctx *Create_index_stmtContext) {}

// EnterDrop_index_stmt is called when production drop_index_stmt is entered.
func (s *BaseCQL3Listener) EnterDrop_index_stmt(ctx *Drop_index_stmtContext) {}

// ExitDrop_index_stmt is called when production drop_index_stmt is exited.
func (s *BaseCQL3Listener) ExitDrop_index_stmt(ctx *Drop_index_stmtContext) {}

// EnterInsert_stmt is called when production insert_stmt is entered.
func (s *BaseCQL3Listener) EnterInsert_stmt(ctx *Insert_stmtContext) {}

// ExitInsert_stmt is called when production insert_stmt is exited.
func (s *BaseCQL3Listener) ExitInsert_stmt(ctx *Insert_stmtContext) {}

// EnterColumn_names is called when production column_names is entered.
func (s *BaseCQL3Listener) EnterColumn_names(ctx *Column_namesContext) {}

// ExitColumn_names is called when production column_names is exited.
func (s *BaseCQL3Listener) ExitColumn_names(ctx *Column_namesContext) {}

// EnterColumn_values is called when production column_values is entered.
func (s *BaseCQL3Listener) EnterColumn_values(ctx *Column_valuesContext) {}

// ExitColumn_values is called when production column_values is exited.
func (s *BaseCQL3Listener) ExitColumn_values(ctx *Column_valuesContext) {}

// EnterUpsert_options is called when production upsert_options is entered.
func (s *BaseCQL3Listener) EnterUpsert_options(ctx *Upsert_optionsContext) {}

// ExitUpsert_options is called when production upsert_options is exited.
func (s *BaseCQL3Listener) ExitUpsert_options(ctx *Upsert_optionsContext) {}

// EnterUpsert_option is called when production upsert_option is entered.
func (s *BaseCQL3Listener) EnterUpsert_option(ctx *Upsert_optionContext) {}

// ExitUpsert_option is called when production upsert_option is exited.
func (s *BaseCQL3Listener) ExitUpsert_option(ctx *Upsert_optionContext) {}

// EnterIndex_name is called when production index_name is entered.
func (s *BaseCQL3Listener) EnterIndex_name(ctx *Index_nameContext) {}

// ExitIndex_name is called when production index_name is exited.
func (s *BaseCQL3Listener) ExitIndex_name(ctx *Index_nameContext) {}

// EnterIndex_class is called when production index_class is entered.
func (s *BaseCQL3Listener) EnterIndex_class(ctx *Index_classContext) {}

// ExitIndex_class is called when production index_class is exited.
func (s *BaseCQL3Listener) ExitIndex_class(ctx *Index_classContext) {}

// EnterIndex_options is called when production index_options is entered.
func (s *BaseCQL3Listener) EnterIndex_options(ctx *Index_optionsContext) {}

// ExitIndex_options is called when production index_options is exited.
func (s *BaseCQL3Listener) ExitIndex_options(ctx *Index_optionsContext) {}

// EnterUpdate_stmt is called when production update_stmt is entered.
func (s *BaseCQL3Listener) EnterUpdate_stmt(ctx *Update_stmtContext) {}

// ExitUpdate_stmt is called when production update_stmt is exited.
func (s *BaseCQL3Listener) ExitUpdate_stmt(ctx *Update_stmtContext) {}

// EnterUpdate_assignments is called when production update_assignments is entered.
func (s *BaseCQL3Listener) EnterUpdate_assignments(ctx *Update_assignmentsContext) {}

// ExitUpdate_assignments is called when production update_assignments is exited.
func (s *BaseCQL3Listener) ExitUpdate_assignments(ctx *Update_assignmentsContext) {}

// EnterUpdate_assignment is called when production update_assignment is entered.
func (s *BaseCQL3Listener) EnterUpdate_assignment(ctx *Update_assignmentContext) {}

// ExitUpdate_assignment is called when production update_assignment is exited.
func (s *BaseCQL3Listener) ExitUpdate_assignment(ctx *Update_assignmentContext) {}

// EnterUpdate_conditions is called when production update_conditions is entered.
func (s *BaseCQL3Listener) EnterUpdate_conditions(ctx *Update_conditionsContext) {}

// ExitUpdate_conditions is called when production update_conditions is exited.
func (s *BaseCQL3Listener) ExitUpdate_conditions(ctx *Update_conditionsContext) {}

// EnterUpdate_condition is called when production update_condition is entered.
func (s *BaseCQL3Listener) EnterUpdate_condition(ctx *Update_conditionContext) {}

// ExitUpdate_condition is called when production update_condition is exited.
func (s *BaseCQL3Listener) ExitUpdate_condition(ctx *Update_conditionContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseCQL3Listener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseCQL3Listener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseCQL3Listener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseCQL3Listener) ExitRelation(ctx *RelationContext) {}

// EnterDelete_stmt is called when production delete_stmt is entered.
func (s *BaseCQL3Listener) EnterDelete_stmt(ctx *Delete_stmtContext) {}

// ExitDelete_stmt is called when production delete_stmt is exited.
func (s *BaseCQL3Listener) ExitDelete_stmt(ctx *Delete_stmtContext) {}

// EnterDelete_conditions is called when production delete_conditions is entered.
func (s *BaseCQL3Listener) EnterDelete_conditions(ctx *Delete_conditionsContext) {}

// ExitDelete_conditions is called when production delete_conditions is exited.
func (s *BaseCQL3Listener) ExitDelete_conditions(ctx *Delete_conditionsContext) {}

// EnterDelete_condition is called when production delete_condition is entered.
func (s *BaseCQL3Listener) EnterDelete_condition(ctx *Delete_conditionContext) {}

// ExitDelete_condition is called when production delete_condition is exited.
func (s *BaseCQL3Listener) ExitDelete_condition(ctx *Delete_conditionContext) {}

// EnterDelete_selections is called when production delete_selections is entered.
func (s *BaseCQL3Listener) EnterDelete_selections(ctx *Delete_selectionsContext) {}

// ExitDelete_selections is called when production delete_selections is exited.
func (s *BaseCQL3Listener) ExitDelete_selections(ctx *Delete_selectionsContext) {}

// EnterDelete_selection is called when production delete_selection is entered.
func (s *BaseCQL3Listener) EnterDelete_selection(ctx *Delete_selectionContext) {}

// ExitDelete_selection is called when production delete_selection is exited.
func (s *BaseCQL3Listener) ExitDelete_selection(ctx *Delete_selectionContext) {}

// EnterBatch_stmt is called when production batch_stmt is entered.
func (s *BaseCQL3Listener) EnterBatch_stmt(ctx *Batch_stmtContext) {}

// ExitBatch_stmt is called when production batch_stmt is exited.
func (s *BaseCQL3Listener) ExitBatch_stmt(ctx *Batch_stmtContext) {}

// EnterBatch_options is called when production batch_options is entered.
func (s *BaseCQL3Listener) EnterBatch_options(ctx *Batch_optionsContext) {}

// ExitBatch_options is called when production batch_options is exited.
func (s *BaseCQL3Listener) ExitBatch_options(ctx *Batch_optionsContext) {}

// EnterBatch_option is called when production batch_option is entered.
func (s *BaseCQL3Listener) EnterBatch_option(ctx *Batch_optionContext) {}

// ExitBatch_option is called when production batch_option is exited.
func (s *BaseCQL3Listener) ExitBatch_option(ctx *Batch_optionContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseCQL3Listener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseCQL3Listener) ExitTable_name(ctx *Table_nameContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BaseCQL3Listener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BaseCQL3Listener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterTable_options is called when production table_options is entered.
func (s *BaseCQL3Listener) EnterTable_options(ctx *Table_optionsContext) {}

// ExitTable_options is called when production table_options is exited.
func (s *BaseCQL3Listener) ExitTable_options(ctx *Table_optionsContext) {}

// EnterTable_option is called when production table_option is entered.
func (s *BaseCQL3Listener) EnterTable_option(ctx *Table_optionContext) {}

// ExitTable_option is called when production table_option is exited.
func (s *BaseCQL3Listener) ExitTable_option(ctx *Table_optionContext) {}

// EnterColumn_definitions is called when production column_definitions is entered.
func (s *BaseCQL3Listener) EnterColumn_definitions(ctx *Column_definitionsContext) {}

// ExitColumn_definitions is called when production column_definitions is exited.
func (s *BaseCQL3Listener) ExitColumn_definitions(ctx *Column_definitionsContext) {}

// EnterColumn_definition is called when production column_definition is entered.
func (s *BaseCQL3Listener) EnterColumn_definition(ctx *Column_definitionContext) {}

// ExitColumn_definition is called when production column_definition is exited.
func (s *BaseCQL3Listener) ExitColumn_definition(ctx *Column_definitionContext) {}

// EnterColumn_type is called when production column_type is entered.
func (s *BaseCQL3Listener) EnterColumn_type(ctx *Column_typeContext) {}

// ExitColumn_type is called when production column_type is exited.
func (s *BaseCQL3Listener) ExitColumn_type(ctx *Column_typeContext) {}

// EnterPrimary_key is called when production primary_key is entered.
func (s *BaseCQL3Listener) EnterPrimary_key(ctx *Primary_keyContext) {}

// ExitPrimary_key is called when production primary_key is exited.
func (s *BaseCQL3Listener) ExitPrimary_key(ctx *Primary_keyContext) {}

// EnterPartition_key is called when production partition_key is entered.
func (s *BaseCQL3Listener) EnterPartition_key(ctx *Partition_keyContext) {}

// ExitPartition_key is called when production partition_key is exited.
func (s *BaseCQL3Listener) ExitPartition_key(ctx *Partition_keyContext) {}

// EnterKeyspace_name is called when production keyspace_name is entered.
func (s *BaseCQL3Listener) EnterKeyspace_name(ctx *Keyspace_nameContext) {}

// ExitKeyspace_name is called when production keyspace_name is exited.
func (s *BaseCQL3Listener) ExitKeyspace_name(ctx *Keyspace_nameContext) {}

// EnterIf_not_exists is called when production if_not_exists is entered.
func (s *BaseCQL3Listener) EnterIf_not_exists(ctx *If_not_existsContext) {}

// ExitIf_not_exists is called when production if_not_exists is exited.
func (s *BaseCQL3Listener) ExitIf_not_exists(ctx *If_not_existsContext) {}

// EnterIf_exists is called when production if_exists is entered.
func (s *BaseCQL3Listener) EnterIf_exists(ctx *If_existsContext) {}

// ExitIf_exists is called when production if_exists is exited.
func (s *BaseCQL3Listener) ExitIf_exists(ctx *If_existsContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCQL3Listener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCQL3Listener) ExitConstant(ctx *ConstantContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseCQL3Listener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseCQL3Listener) ExitVariable(ctx *VariableContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseCQL3Listener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseCQL3Listener) ExitTerm(ctx *TermContext) {}

// EnterCollection is called when production collection is entered.
func (s *BaseCQL3Listener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BaseCQL3Listener) ExitCollection(ctx *CollectionContext) {}

// EnterMap is called when production map is entered.
func (s *BaseCQL3Listener) EnterMap(ctx *MapContext) {}

// ExitMap is called when production map is exited.
func (s *BaseCQL3Listener) ExitMap(ctx *MapContext) {}

// EnterSet is called when production set is entered.
func (s *BaseCQL3Listener) EnterSet(ctx *SetContext) {}

// ExitSet is called when production set is exited.
func (s *BaseCQL3Listener) ExitSet(ctx *SetContext) {}

// EnterList is called when production list is entered.
func (s *BaseCQL3Listener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseCQL3Listener) ExitList(ctx *ListContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseCQL3Listener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseCQL3Listener) ExitFunction(ctx *FunctionContext) {}

// EnterProperties is called when production properties is entered.
func (s *BaseCQL3Listener) EnterProperties(ctx *PropertiesContext) {}

// ExitProperties is called when production properties is exited.
func (s *BaseCQL3Listener) ExitProperties(ctx *PropertiesContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseCQL3Listener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseCQL3Listener) ExitProperty(ctx *PropertyContext) {}

// EnterProperty_name is called when production property_name is entered.
func (s *BaseCQL3Listener) EnterProperty_name(ctx *Property_nameContext) {}

// ExitProperty_name is called when production property_name is exited.
func (s *BaseCQL3Listener) ExitProperty_name(ctx *Property_nameContext) {}

// EnterProperty_value is called when production property_value is entered.
func (s *BaseCQL3Listener) EnterProperty_value(ctx *Property_valueContext) {}

// ExitProperty_value is called when production property_value is exited.
func (s *BaseCQL3Listener) ExitProperty_value(ctx *Property_valueContext) {}

// EnterData_type is called when production data_type is entered.
func (s *BaseCQL3Listener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BaseCQL3Listener) ExitData_type(ctx *Data_typeContext) {}

// EnterNative_type is called when production native_type is entered.
func (s *BaseCQL3Listener) EnterNative_type(ctx *Native_typeContext) {}

// ExitNative_type is called when production native_type is exited.
func (s *BaseCQL3Listener) ExitNative_type(ctx *Native_typeContext) {}

// EnterCollection_type is called when production collection_type is entered.
func (s *BaseCQL3Listener) EnterCollection_type(ctx *Collection_typeContext) {}

// ExitCollection_type is called when production collection_type is exited.
func (s *BaseCQL3Listener) ExitCollection_type(ctx *Collection_typeContext) {}

// EnterBool is called when production bool is entered.
func (s *BaseCQL3Listener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BaseCQL3Listener) ExitBool(ctx *BoolContext) {}
