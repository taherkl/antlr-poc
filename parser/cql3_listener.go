// Generated from CQL3.g4 by ANTLR 4.7.

package parser // CQL3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CQL3Listener is a complete listener for a parse tree produced by CQL3Parser.
type CQL3Listener interface {
	antlr.ParseTreeListener

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDml_statements is called when entering the dml_statements production.
	EnterDml_statements(c *Dml_statementsContext)

	// EnterDml_statement is called when entering the dml_statement production.
	EnterDml_statement(c *Dml_statementContext)

	// EnterCreate_keyspace_stmt is called when entering the create_keyspace_stmt production.
	EnterCreate_keyspace_stmt(c *Create_keyspace_stmtContext)

	// EnterAlter_keyspace_stmt is called when entering the alter_keyspace_stmt production.
	EnterAlter_keyspace_stmt(c *Alter_keyspace_stmtContext)

	// EnterDrop_keyspace_stmt is called when entering the drop_keyspace_stmt production.
	EnterDrop_keyspace_stmt(c *Drop_keyspace_stmtContext)

	// EnterUse_stmt is called when entering the use_stmt production.
	EnterUse_stmt(c *Use_stmtContext)

	// EnterCreate_table_stmt is called when entering the create_table_stmt production.
	EnterCreate_table_stmt(c *Create_table_stmtContext)

	// EnterAlter_table_stmt is called when entering the alter_table_stmt production.
	EnterAlter_table_stmt(c *Alter_table_stmtContext)

	// EnterAlter_table_instruction is called when entering the alter_table_instruction production.
	EnterAlter_table_instruction(c *Alter_table_instructionContext)

	// EnterDrop_table_stmt is called when entering the drop_table_stmt production.
	EnterDrop_table_stmt(c *Drop_table_stmtContext)

	// EnterTruncate_table_stmt is called when entering the truncate_table_stmt production.
	EnterTruncate_table_stmt(c *Truncate_table_stmtContext)

	// EnterCreate_index_stmt is called when entering the create_index_stmt production.
	EnterCreate_index_stmt(c *Create_index_stmtContext)

	// EnterDrop_index_stmt is called when entering the drop_index_stmt production.
	EnterDrop_index_stmt(c *Drop_index_stmtContext)

	// EnterInsert_stmt is called when entering the insert_stmt production.
	EnterInsert_stmt(c *Insert_stmtContext)

	// EnterColumn_names is called when entering the column_names production.
	EnterColumn_names(c *Column_namesContext)

	// EnterColumn_values is called when entering the column_values production.
	EnterColumn_values(c *Column_valuesContext)

	// EnterUpsert_options is called when entering the upsert_options production.
	EnterUpsert_options(c *Upsert_optionsContext)

	// EnterUpsert_option is called when entering the upsert_option production.
	EnterUpsert_option(c *Upsert_optionContext)

	// EnterIndex_name is called when entering the index_name production.
	EnterIndex_name(c *Index_nameContext)

	// EnterIndex_class is called when entering the index_class production.
	EnterIndex_class(c *Index_classContext)

	// EnterIndex_options is called when entering the index_options production.
	EnterIndex_options(c *Index_optionsContext)

	// EnterUpdate_stmt is called when entering the update_stmt production.
	EnterUpdate_stmt(c *Update_stmtContext)

	// EnterUpdate_assignments is called when entering the update_assignments production.
	EnterUpdate_assignments(c *Update_assignmentsContext)

	// EnterUpdate_assignment is called when entering the update_assignment production.
	EnterUpdate_assignment(c *Update_assignmentContext)

	// EnterUpdate_conditions is called when entering the update_conditions production.
	EnterUpdate_conditions(c *Update_conditionsContext)

	// EnterUpdate_condition is called when entering the update_condition production.
	EnterUpdate_condition(c *Update_conditionContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterDelete_stmt is called when entering the delete_stmt production.
	EnterDelete_stmt(c *Delete_stmtContext)

	// EnterDelete_conditions is called when entering the delete_conditions production.
	EnterDelete_conditions(c *Delete_conditionsContext)

	// EnterDelete_condition is called when entering the delete_condition production.
	EnterDelete_condition(c *Delete_conditionContext)

	// EnterDelete_selections is called when entering the delete_selections production.
	EnterDelete_selections(c *Delete_selectionsContext)

	// EnterDelete_selection is called when entering the delete_selection production.
	EnterDelete_selection(c *Delete_selectionContext)

	// EnterBatch_stmt is called when entering the batch_stmt production.
	EnterBatch_stmt(c *Batch_stmtContext)

	// EnterBatch_options is called when entering the batch_options production.
	EnterBatch_options(c *Batch_optionsContext)

	// EnterBatch_option is called when entering the batch_option production.
	EnterBatch_option(c *Batch_optionContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterTable_options is called when entering the table_options production.
	EnterTable_options(c *Table_optionsContext)

	// EnterTable_option is called when entering the table_option production.
	EnterTable_option(c *Table_optionContext)

	// EnterColumn_definitions is called when entering the column_definitions production.
	EnterColumn_definitions(c *Column_definitionsContext)

	// EnterColumn_definition is called when entering the column_definition production.
	EnterColumn_definition(c *Column_definitionContext)

	// EnterColumn_type is called when entering the column_type production.
	EnterColumn_type(c *Column_typeContext)

	// EnterPrimary_key is called when entering the primary_key production.
	EnterPrimary_key(c *Primary_keyContext)

	// EnterPartition_key is called when entering the partition_key production.
	EnterPartition_key(c *Partition_keyContext)

	// EnterKeyspace_name is called when entering the keyspace_name production.
	EnterKeyspace_name(c *Keyspace_nameContext)

	// EnterIf_not_exists is called when entering the if_not_exists production.
	EnterIf_not_exists(c *If_not_existsContext)

	// EnterIf_exists is called when entering the if_exists production.
	EnterIf_exists(c *If_existsContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterCollection is called when entering the collection production.
	EnterCollection(c *CollectionContext)

	// EnterMap is called when entering the map production.
	EnterMap(c *MapContext)

	// EnterSet is called when entering the set production.
	EnterSet(c *SetContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterProperties is called when entering the properties production.
	EnterProperties(c *PropertiesContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterProperty_name is called when entering the property_name production.
	EnterProperty_name(c *Property_nameContext)

	// EnterProperty_value is called when entering the property_value production.
	EnterProperty_value(c *Property_valueContext)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterNative_type is called when entering the native_type production.
	EnterNative_type(c *Native_typeContext)

	// EnterCollection_type is called when entering the collection_type production.
	EnterCollection_type(c *Collection_typeContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDml_statements is called when exiting the dml_statements production.
	ExitDml_statements(c *Dml_statementsContext)

	// ExitDml_statement is called when exiting the dml_statement production.
	ExitDml_statement(c *Dml_statementContext)

	// ExitCreate_keyspace_stmt is called when exiting the create_keyspace_stmt production.
	ExitCreate_keyspace_stmt(c *Create_keyspace_stmtContext)

	// ExitAlter_keyspace_stmt is called when exiting the alter_keyspace_stmt production.
	ExitAlter_keyspace_stmt(c *Alter_keyspace_stmtContext)

	// ExitDrop_keyspace_stmt is called when exiting the drop_keyspace_stmt production.
	ExitDrop_keyspace_stmt(c *Drop_keyspace_stmtContext)

	// ExitUse_stmt is called when exiting the use_stmt production.
	ExitUse_stmt(c *Use_stmtContext)

	// ExitCreate_table_stmt is called when exiting the create_table_stmt production.
	ExitCreate_table_stmt(c *Create_table_stmtContext)

	// ExitAlter_table_stmt is called when exiting the alter_table_stmt production.
	ExitAlter_table_stmt(c *Alter_table_stmtContext)

	// ExitAlter_table_instruction is called when exiting the alter_table_instruction production.
	ExitAlter_table_instruction(c *Alter_table_instructionContext)

	// ExitDrop_table_stmt is called when exiting the drop_table_stmt production.
	ExitDrop_table_stmt(c *Drop_table_stmtContext)

	// ExitTruncate_table_stmt is called when exiting the truncate_table_stmt production.
	ExitTruncate_table_stmt(c *Truncate_table_stmtContext)

	// ExitCreate_index_stmt is called when exiting the create_index_stmt production.
	ExitCreate_index_stmt(c *Create_index_stmtContext)

	// ExitDrop_index_stmt is called when exiting the drop_index_stmt production.
	ExitDrop_index_stmt(c *Drop_index_stmtContext)

	// ExitInsert_stmt is called when exiting the insert_stmt production.
	ExitInsert_stmt(c *Insert_stmtContext)

	// ExitColumn_names is called when exiting the column_names production.
	ExitColumn_names(c *Column_namesContext)

	// ExitColumn_values is called when exiting the column_values production.
	ExitColumn_values(c *Column_valuesContext)

	// ExitUpsert_options is called when exiting the upsert_options production.
	ExitUpsert_options(c *Upsert_optionsContext)

	// ExitUpsert_option is called when exiting the upsert_option production.
	ExitUpsert_option(c *Upsert_optionContext)

	// ExitIndex_name is called when exiting the index_name production.
	ExitIndex_name(c *Index_nameContext)

	// ExitIndex_class is called when exiting the index_class production.
	ExitIndex_class(c *Index_classContext)

	// ExitIndex_options is called when exiting the index_options production.
	ExitIndex_options(c *Index_optionsContext)

	// ExitUpdate_stmt is called when exiting the update_stmt production.
	ExitUpdate_stmt(c *Update_stmtContext)

	// ExitUpdate_assignments is called when exiting the update_assignments production.
	ExitUpdate_assignments(c *Update_assignmentsContext)

	// ExitUpdate_assignment is called when exiting the update_assignment production.
	ExitUpdate_assignment(c *Update_assignmentContext)

	// ExitUpdate_conditions is called when exiting the update_conditions production.
	ExitUpdate_conditions(c *Update_conditionsContext)

	// ExitUpdate_condition is called when exiting the update_condition production.
	ExitUpdate_condition(c *Update_conditionContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitDelete_stmt is called when exiting the delete_stmt production.
	ExitDelete_stmt(c *Delete_stmtContext)

	// ExitDelete_conditions is called when exiting the delete_conditions production.
	ExitDelete_conditions(c *Delete_conditionsContext)

	// ExitDelete_condition is called when exiting the delete_condition production.
	ExitDelete_condition(c *Delete_conditionContext)

	// ExitDelete_selections is called when exiting the delete_selections production.
	ExitDelete_selections(c *Delete_selectionsContext)

	// ExitDelete_selection is called when exiting the delete_selection production.
	ExitDelete_selection(c *Delete_selectionContext)

	// ExitBatch_stmt is called when exiting the batch_stmt production.
	ExitBatch_stmt(c *Batch_stmtContext)

	// ExitBatch_options is called when exiting the batch_options production.
	ExitBatch_options(c *Batch_optionsContext)

	// ExitBatch_option is called when exiting the batch_option production.
	ExitBatch_option(c *Batch_optionContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitTable_options is called when exiting the table_options production.
	ExitTable_options(c *Table_optionsContext)

	// ExitTable_option is called when exiting the table_option production.
	ExitTable_option(c *Table_optionContext)

	// ExitColumn_definitions is called when exiting the column_definitions production.
	ExitColumn_definitions(c *Column_definitionsContext)

	// ExitColumn_definition is called when exiting the column_definition production.
	ExitColumn_definition(c *Column_definitionContext)

	// ExitColumn_type is called when exiting the column_type production.
	ExitColumn_type(c *Column_typeContext)

	// ExitPrimary_key is called when exiting the primary_key production.
	ExitPrimary_key(c *Primary_keyContext)

	// ExitPartition_key is called when exiting the partition_key production.
	ExitPartition_key(c *Partition_keyContext)

	// ExitKeyspace_name is called when exiting the keyspace_name production.
	ExitKeyspace_name(c *Keyspace_nameContext)

	// ExitIf_not_exists is called when exiting the if_not_exists production.
	ExitIf_not_exists(c *If_not_existsContext)

	// ExitIf_exists is called when exiting the if_exists production.
	ExitIf_exists(c *If_existsContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitCollection is called when exiting the collection production.
	ExitCollection(c *CollectionContext)

	// ExitMap is called when exiting the map production.
	ExitMap(c *MapContext)

	// ExitSet is called when exiting the set production.
	ExitSet(c *SetContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitProperties is called when exiting the properties production.
	ExitProperties(c *PropertiesContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitProperty_name is called when exiting the property_name production.
	ExitProperty_name(c *Property_nameContext)

	// ExitProperty_value is called when exiting the property_value production.
	ExitProperty_value(c *Property_valueContext)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitNative_type is called when exiting the native_type production.
	ExitNative_type(c *Native_typeContext)

	// ExitCollection_type is called when exiting the collection_type production.
	ExitCollection_type(c *Collection_typeContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)
}
